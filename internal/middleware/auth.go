package middleware

import (
	"crypto/ed25519"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/blake2b"
)

// AuthMiddleware contains the authentication middleware functions
type AuthMiddleware struct{}

// NewAuthMiddleware creates a new instance of AuthMiddleware
func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

// getAuthHeader generates an authorization header with signature
func getAuthHeader(payload []byte) (string, error) {
	var authHeader string

	privateKey := os.Getenv("PRIVATE_KEY")
	currentTime := int(time.Now().Unix())

	//ttl we are using is 3600 seconds (1 hour)
	ttl := 3600

	signature, err := signRequest(privateKey, payload, currentTime, ttl)
	if err != nil {
		fmt.Println("Could not compute signature", err)
		return authHeader, err
	}

	subscriberID := os.Getenv("SUBSCRIBER_ID")
	if subscriberID == "" {
		subscriberID = "buyer-app.ondc.org"
	}

	uniqueKeyID := os.Getenv("UNIQUE_KEY_ID")
	if uniqueKeyID == "" {
		uniqueKeyID = "207"
	}
	authHeader = fmt.Sprintf(`Signature keyId="%s|%s|ed25519",algorithm="ed25519",created="%d",expires="%d",headers="(created) (expires) digest",signature="%s"`, subscriberID, uniqueKeyID, currentTime, currentTime+ttl, signature)

	return authHeader, nil
}

// verifyRequest verifies the authorization header
func verifyRequest(authHeader string, payload []byte) bool {
	publicKey := os.Getenv("PUBLIC_KEY")

	_, created, expires, signature, err := parseAuthHeader(authHeader)
	if err != nil {
		return false
	}

	//compute blake 512 hash over the payload
	hash := blake2b.Sum512(payload)
	digest := base64Encode(hash[:])

	//create a signature and then sign with ed25519 private key
	computedMessage := fmt.Sprintf("(created): %s\n(expires): %s\ndigest: BLAKE-512=%s", created, expires, digest)
	publicKeyBytes, err := base64Decode(publicKey)
	if err != nil {
		fmt.Println("Error decoding public key", err)
		return false
	}
	receivedSignature, err := base64Decode(signature)
	if err != nil {
		fmt.Println("Unable to base64 decode received signature", err)
		return false
	}
	return ed25519.Verify(publicKeyBytes, []byte(computedMessage), receivedSignature)
}

// parseAuthHeader parses the authorization header
func parseAuthHeader(authHeader string) (string, string, string, string, error) {
	signatureRegex := regexp.MustCompile(`keyId=\"(.+?)\".+?created=\"(.+?)\".+?expires=\"(.+?)\".+?signature=\"(.+?)\"`)
	groups := signatureRegex.FindAllStringSubmatch(authHeader, -1)
	if len(groups) > 0 && len(groups[0]) > 4 {
		return groups[0][1], groups[0][2], groups[0][3], groups[0][4], nil
	}
	fmt.Println("Error parsing auth header. Please make sure that the auth headers passed as command line argument is valid")
	return "", "", "", "", errors.New("error parsing auth header")
}

// signRequest signs the request payload
func signRequest(privateKey string, payload []byte, currentTime int, ttl int) (string, error) {
	//compute blake 512 hash over the payload
	hash := blake2b.Sum512(payload)
	digest := base64Encode(hash[:])

	//create a signature and then sign with ed25519 private key
	created := strconv.Itoa(currentTime)
	expires := strconv.Itoa(currentTime + ttl)
	signingString := fmt.Sprintf("(created): %s\n(expires): %s\ndigest: BLAKE-512=%s", created, expires, digest)

	privateKeyBytes, err := base64Decode(privateKey)
	if err != nil {
		return "", fmt.Errorf("error decoding private key: %v", err)
	}

	signature := ed25519.Sign(privateKeyBytes, []byte(signingString))
	return base64Encode(signature), nil
}

// base64Encode encodes bytes to base64 string
func base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// base64Decode decodes base64 string to bytes
func base64Decode(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
}

// GenerateAuthHeader middleware for outgoing requests
func (m *AuthMiddleware) GenerateAuthHeader() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the request body for signature generation
		body := c.Body()

		// Generate auth header
		authHeader, err := getAuthHeader(body)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to generate auth header",
			})
		}

		// Set the authorization header
		c.Set("Authorization", authHeader)

		return c.Next()
	}
}

// VerifyAuthHeader middleware for incoming requests
func (m *AuthMiddleware) VerifyAuthHeader() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing authorization header",
			})
		}

		// Read the request body
		body := c.Body()

		// Verify the request
		if !verifyRequest(authHeader, body) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid signature",
			})
		}

		return c.Next()
	}
}

// Optional: Middleware that only verifies if auth header is present
func (m *AuthMiddleware) OptionalVerifyAuthHeader() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the authorization header
		authHeader := c.Get("Authorization")

		// If no auth header is present, skip verification
		if authHeader == "" {
			return c.Next()
		}

		// Read the request body
		body := c.Body()

		// Verify the request
		if !verifyRequest(authHeader, body) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid signature",
			})
		}

		return c.Next()
	}
}

// AddAuthHeaderToRequest adds auth header to outgoing HTTP requests
func AddAuthHeaderToRequest(payload []byte) (map[string]string, error) {
	authHeader, err := getAuthHeader(payload)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Authorization": authHeader,
		"Content-Type":  "application/json",
	}

	return headers, nil
}

// ValidateIncomingRequest validates incoming request with auth header
func ValidateIncomingRequest(authHeader string, payload []byte) error {
	if authHeader == "" {
		return errors.New("missing authorization header")
	}

	if !verifyRequest(authHeader, payload) {
		return errors.New("invalid signature")
	}

	return nil
}

// Helper function to read request body without consuming it
func getRequestBody(c *fiber.Ctx) ([]byte, error) {
	// Get body
	body := c.Body()

	// Create a copy since body might be consumed
	bodyCopy := make([]byte, len(body))
	copy(bodyCopy, body)

	return bodyCopy, nil
}

// Global middleware instances
var DefaultAuthMiddleware = NewAuthMiddleware()

// Convenience functions
func GenerateAuthHeaderMiddleware() fiber.Handler {
	return DefaultAuthMiddleware.GenerateAuthHeader()
}

func VerifyAuthHeaderMiddleware() fiber.Handler {
	return DefaultAuthMiddleware.VerifyAuthHeader()
}

func OptionalVerifyAuthHeaderMiddleware() fiber.Handler {
	return DefaultAuthMiddleware.OptionalVerifyAuthHeader()
}
