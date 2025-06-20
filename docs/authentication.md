# ONDC Authentication Middleware

This document explains how to use the ONDC authentication middleware for generating and verifying ED25519 signatures with BLAKE-512 hashing.

## Overview

The authentication middleware automatically:
- **Adds authentication headers** to all outgoing buyer action API requests
- **Verifies authentication headers** on incoming callback requests (optional)
- **Supports ONDC v1.1.0** signature standards with ED25519 + BLAKE-512

## Environment Variables

Set the following environment variables in your `.env` file:

```bash
# Required for signature generation
PRIVATE_KEY=your_base64_encoded_private_key
PUBLIC_KEY=your_base64_encoded_public_key

# Optional - defaults provided
SUBSCRIBER_ID=buyer-app.ondc.org
UNIQUE_KEY_ID=207
```

## How It Works

### 1. Outgoing Requests (Buyer Action APIs)

Authentication headers are **automatically added** to all outgoing requests in the `OndcService`:

```go
// In internal/services/ondc_service.go
func (s *OndcService) forwardRequest(endpoint string, req interface{}) error {
    jsonData, _ := json.Marshal(req)
    
    // Auth headers are automatically added here
    headers, err := middleware.AddAuthHeaderToRequest(jsonData)
    if err != nil {
        log.Printf("Error generating auth headers: %v", err)
    } else {
        for key, value := range headers {
            httpReq.Header.Set(key, value)
        }
    }
    // ... rest of the request
}
```

**Supported endpoints with auto-auth:**
- `/v1/buyer/search`
- `/v1/buyer/select` 
- `/v1/buyer/init`
- `/v1/buyer/confirm`
- `/v1/buyer/status`
- `/v1/buyer/cancel`
- `/v1/buyer/update`
- `/v1/buyer/track`

### 2. Incoming Requests (Callback APIs)

Callback endpoints use **optional verification** - they verify auth headers if present:

```go
// In internal/api/routes/routes.go
buyerCallbacks := buyer.Group("", authMiddleware.OptionalVerifyAuthHeader())
buyerCallbacks.Post("/on_search", callbackHandlers.HandleOnSearch)
// ... other callbacks
```

**Protected callback endpoints:**
- `/v1/buyer/on_search`
- `/v1/buyer/on_select`
- `/v1/buyer/on_init`
- `/v1/buyer/on_confirm`
- `/v1/buyer/on_status`
- `/v1/buyer/on_cancel`
- `/v1/buyer/on_update`
- `/v1/buyer/on_track`

## Generated Headers

The middleware generates ONDC-compliant authorization headers:

```
Authorization: Signature keyId="buyer-app.ondc.org|207|ed25519",algorithm="ed25519",created="1640995200",expires="1640998800",headers="(created) (expires) digest",signature="base64_signature"
Content-Type: application/json
```

## Usage Examples

### Manual Header Generation

```go
import "ondc-buyer-app/internal/middleware"

// For manual requests
payload := []byte(`{"your": "json_data"}`)
headers, err := middleware.AddAuthHeaderToRequest(payload)
if err != nil {
    return err
}

// headers now contains Authorization and Content-Type
```

### Manual Verification

```go
import "ondc-buyer-app/internal/middleware"

// Verify incoming request manually
authHeader := c.Get("Authorization")
payload := c.Body()

err := middleware.ValidateIncomingRequest(authHeader, payload)
if err != nil {
    return c.Status(401).JSON(fiber.Map{"error": err.Error()})
}
```

### Custom Middleware Usage

```go
// For strict verification (required auth header)
app.Use("/protected", middleware.VerifyAuthHeaderMiddleware())

// For optional verification (verify if present)
app.Use("/optional", middleware.OptionalVerifyAuthHeaderMiddleware())

// For adding auth to outgoing requests
app.Use("/outgoing", middleware.GenerateAuthHeaderMiddleware())
```

## Testing

Test your authentication setup:

```bash
# Test outgoing request (should add auth headers automatically)
curl -X POST http://localhost:4000/v1/buyer/search \
  -H "Content-Type: application/json" \
  -d '{"context": {"domain": "ONDC:TRV10"}}'

# Test callback (auth header optional)
curl -X POST http://localhost:4000/v1/buyer/on_search \
  -H "Content-Type: application/json" \
  -d '{"message": {"catalog": {}}}'
```

## Key Generation

Generate ED25519 keys for testing:

```bash
# Using OpenSSL
openssl genpkey -algorithm Ed25519 -out private_key.pem
openssl pkey -in private_key.pem -pubout -out public_key.pem

# Convert to base64 for environment variables
openssl pkey -in private_key.pem -raw -outform DER | base64
openssl pkey -pubin -in public_key.pem -raw -outform DER | base64
```

## Security Notes

- **Private keys** should be kept secure and not committed to version control
- **TTL** is set to 3600 seconds (1 hour) for signature validity
- **BLAKE-512** hashing ensures request integrity
- **ED25519** provides fast signature verification
- Headers are automatically logged (but not the signature content for security)

## Troubleshooting

1. **"Error generating auth headers"**: Check PRIVATE_KEY environment variable
2. **"Invalid signature"**: Verify PUBLIC_KEY matches the signing private key
3. **"Missing authorization header"**: Check if middleware is correctly applied
4. **Timeout errors**: Ensure BPP_URI is correct and reachable

## Integration Status

✅ **Automatic auth headers** added to all buyer action APIs  
✅ **Optional verification** on callback endpoints  
✅ **ONDC v1.1.0 compliant** signatures  
✅ **Environment-based configuration**  
✅ **Comprehensive logging** for debugging 