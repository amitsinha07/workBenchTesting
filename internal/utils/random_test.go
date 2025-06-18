package utils

import (
	"strconv"
	"strings"
	"testing"
)

func TestGenerateRandomNumber(t *testing.T) {
	// Test different lengths
	lengths := []int{5, 10, 15}

	for _, length := range lengths {
		number := GenerateRandomNumber(length)

		// Check if length is correct
		if len(number) != length {
			t.Errorf("Expected length %d, got %d for number: %s", length, len(number), number)
		}

		// Check if it's a valid number
		if _, err := strconv.ParseInt(number, 10, 64); err != nil {
			t.Errorf("Generated number is not valid: %s, error: %v", number, err)
		}

		// Check if it doesn't start with 0 (except for single digit)
		if length > 1 && number[0] == '0' {
			t.Errorf("Multi-digit number should not start with 0: %s", number)
		}
	}
}

func TestGeneratePhoneNumber(t *testing.T) {
	phone := GeneratePhoneNumber()

	// Check length
	if len(phone) != 10 {
		t.Errorf("Expected phone number length 10, got %d: %s", len(phone), phone)
	}

	// Check if it's numeric
	if _, err := strconv.ParseInt(phone, 10, 64); err != nil {
		t.Errorf("Phone number is not numeric: %s, error: %v", phone, err)
	}
}

func TestGenerateTransactionID(t *testing.T) {
	txnID := GenerateTransactionID()

	// Check if it starts with "txn-"
	if !strings.HasPrefix(txnID, "txn-") {
		t.Errorf("Transaction ID should start with 'txn-': %s", txnID)
	}

	// Check if it has the expected format: txn-timestamp-randomnumber
	parts := strings.Split(txnID, "-")
	if len(parts) != 3 {
		t.Errorf("Transaction ID should have 3 parts separated by '-': %s", txnID)
	}
}

func TestGenerateMessageID(t *testing.T) {
	msgID := GenerateMessageID()

	// Check if it starts with "msg-"
	if !strings.HasPrefix(msgID, "msg-") {
		t.Errorf("Message ID should start with 'msg-': %s", msgID)
	}

	// Check if it has the expected format: msg-timestamp-randomnumber
	parts := strings.Split(msgID, "-")
	if len(parts) != 3 {
		t.Errorf("Message ID should have 3 parts separated by '-': %s", msgID)
	}
}

func TestGenerateOrderID(t *testing.T) {
	orderID := GenerateOrderID()

	// Check if it starts with "order-"
	if !strings.HasPrefix(orderID, "order-") {
		t.Errorf("Order ID should start with 'order-': %s", orderID)
	}

	// Check if the random part is 8 digits
	parts := strings.Split(orderID, "-")
	if len(parts) != 2 || len(parts[1]) != 8 {
		t.Errorf("Order ID should have format 'order-xxxxxxxx': %s", orderID)
	}
}

func TestGenerateACKNumber(t *testing.T) {
	ackNumber := GenerateACKNumber()

	// Check length
	if len(ackNumber) != 10 {
		t.Errorf("Expected ACK number length 10, got %d: %s", len(ackNumber), ackNumber)
	}

	// Check if it's numeric
	if _, err := strconv.ParseInt(ackNumber, 10, 64); err != nil {
		t.Errorf("ACK number is not numeric: %s, error: %v", ackNumber, err)
	}
}

func TestGenerateRandomPrice(t *testing.T) {
	price := GenerateRandomPrice(10.0, 100.0)

	// Check if it's a valid float
	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		t.Errorf("Generated price is not valid: %s, error: %v", price, err)
	}

	// Check if it's within range
	if priceFloat < 10.0 || priceFloat > 100.0 {
		t.Errorf("Price should be between 10.0 and 100.0, got: %f", priceFloat)
	}

	// Check if it has 2 decimal places
	if !strings.Contains(price, ".") {
		t.Errorf("Price should have decimal places: %s", price)
	}

	parts := strings.Split(price, ".")
	if len(parts) != 2 || len(parts[1]) != 2 {
		t.Errorf("Price should have exactly 2 decimal places: %s", price)
	}
}

func TestUniqueness(t *testing.T) {
	// Test that multiple calls generate different numbers
	numbers := make(map[string]bool)

	for i := 0; i < 100; i++ {
		number := GenerateRandomNumber(10)
		if numbers[number] {
			t.Errorf("Generated duplicate number: %s", number)
		}
		numbers[number] = true
	}
}

func BenchmarkGenerateRandomNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateRandomNumber(10)
	}
}

func BenchmarkGenerateACKNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateACKNumber()
	}
}
