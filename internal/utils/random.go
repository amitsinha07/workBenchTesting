package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"time"
)

// RandomGenerator provides methods for generating random values
type RandomGenerator struct{}

// NewRandomGenerator creates a new instance of RandomGenerator
func NewRandomGenerator() *RandomGenerator {
	return &RandomGenerator{}
}

// GenerateRandomNumber generates a random number string of specified length
func (r *RandomGenerator) GenerateRandomNumber(length int) string {
	if length <= 0 {
		length = 10 // default length
	}

	// Generate random number between 10^(length-1) and 10^length - 1
	min := int64(1)
	for i := 1; i < length; i++ {
		min *= 10
	}
	max := min*10 - 1

	// Use crypto/rand for better randomness
	n, err := rand.Int(rand.Reader, big.NewInt(max-min+1))
	if err != nil {
		// Fallback to timestamp-based generation
		return r.generateFallbackNumber(length)
	}

	return strconv.FormatInt(n.Int64()+min, 10)
}

// GeneratePhoneNumber generates a random 10-digit phone number
func (r *RandomGenerator) GeneratePhoneNumber() string {
	return r.GenerateRandomNumber(10)
}

// GenerateTransactionID generates a random transaction ID
func (r *RandomGenerator) GenerateTransactionID() string {
	return fmt.Sprintf("txn-%d-%s", time.Now().Unix(), r.GenerateRandomNumber(6))
}

// GenerateMessageID generates a random message ID
func (r *RandomGenerator) GenerateMessageID() string {
	return fmt.Sprintf("msg-%d-%s", time.Now().Unix(), r.GenerateRandomNumber(6))
}

// GenerateOrderID generates a random order ID
func (r *RandomGenerator) GenerateOrderID() string {
	return fmt.Sprintf("order-%s", r.GenerateRandomNumber(8))
}

// GenerateACKNumber generates a random ACK number for ONDC responses
func (r *RandomGenerator) GenerateACKNumber() string {
	return r.GenerateRandomNumber(10)
}

// GenerateRandomFloat generates a random float between min and max
func (r *RandomGenerator) GenerateRandomFloat(min, max float64) float64 {
	if min >= max {
		return min
	}

	// Generate random number between 0 and (max-min)
	range64 := int64((max - min) * 100) // multiply by 100 for 2 decimal precision
	n, err := rand.Int(rand.Reader, big.NewInt(range64))
	if err != nil {
		return min
	}

	return min + float64(n.Int64())/100.0
}

// GenerateRandomPrice generates a random price between min and max
func (r *RandomGenerator) GenerateRandomPrice(min, max float64) string {
	price := r.GenerateRandomFloat(min, max)
	return fmt.Sprintf("%.2f", price)
}

// generateFallbackNumber generates a number using timestamp as fallback
func (r *RandomGenerator) generateFallbackNumber(length int) string {
	timestamp := time.Now().UnixNano()
	timestampStr := strconv.FormatInt(timestamp, 10)

	// Take last 'length' digits
	if len(timestampStr) >= length {
		return timestampStr[len(timestampStr)-length:]
	}

	// Pad with zeros if needed
	for len(timestampStr) < length {
		timestampStr = "0" + timestampStr
	}

	return timestampStr
}

// Global instance for convenience
var DefaultRandomGenerator = NewRandomGenerator()

// Convenience functions using the default generator
func GenerateRandomNumber(length int) string {
	return DefaultRandomGenerator.GenerateRandomNumber(length)
}

func GeneratePhoneNumber() string {
	return DefaultRandomGenerator.GeneratePhoneNumber()
}

func GenerateTransactionID() string {
	return DefaultRandomGenerator.GenerateTransactionID()
}

func GenerateMessageID() string {
	return DefaultRandomGenerator.GenerateMessageID()
}

func GenerateOrderID() string {
	return DefaultRandomGenerator.GenerateOrderID()
}

func GenerateACKNumber() string {
	return DefaultRandomGenerator.GenerateACKNumber()
}

func GenerateRandomPrice(min, max float64) string {
	return DefaultRandomGenerator.GenerateRandomPrice(min, max)
}
