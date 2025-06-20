package services

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"ondc-buyer-app/internal/config"
	"ondc-buyer-app/internal/middleware"
	"strings"
	"time"
)

// OndcService handles communication with ONDC network
type OndcService struct {
	config     *config.Config
	httpClient *http.Client
}

// NewOndcService creates a new ONDC service instance
func NewOndcService(cfg *config.Config) *OndcService {
	return &OndcService{
		config: cfg,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// maskSignature masks the signature part of authorization header for security
func maskSignature(authHeader string) string {
	if strings.Contains(authHeader, "signature=") {
		parts := strings.Split(authHeader, "signature=")
		if len(parts) > 1 {
			// Show first part + masked signature
			return parts[0] + "signature=***MASKED***"
		}
	}
	return authHeader
}

// forwardRequest is a helper function to forward requests to BPP with auth headers
func (s *OndcService) forwardRequest(endpoint string, req interface{}) error {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	// Construct the full URL
	fullURL := s.config.BppURI + endpoint

	// Log the BPP URI before forwarding
	log.Printf("Forwarding %s request to BPP URI: %s", endpoint, fullURL)

	// Create HTTP request
	httpReq, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error creating request to %s: %v", fullURL, err)
		return err
	}

	// Add authentication headers using the middleware function
	headers, err := middleware.AddAuthHeaderToRequest(jsonData)
	if err != nil {
		log.Printf("Error generating auth headers for %s: %v", endpoint, err)
		// Continue without auth headers - you may want to make this more strict
	} else {
		// Add all headers from the middleware
		for key, value := range headers {
			httpReq.Header.Set(key, value)
		}
		log.Printf("Added authentication headers for request to %s", endpoint)
	}

	// Log complete request details
	log.Printf("=== OUTGOING REQUEST DETAILS ===")
	log.Printf("Method: %s", httpReq.Method)
	log.Printf("URL: %s", httpReq.URL.String())
	log.Printf("Headers:")
	for name, values := range httpReq.Header {
		for _, value := range values {
			// Don't log the full signature for security, just show it exists
			if name == "Authorization" {
				log.Printf("  %s: %s", name, maskSignature(value))
			} else {
				log.Printf("  %s: %s", name, value)
			}
		}
	}
	log.Printf("Request Body: %s", string(jsonData))
	log.Printf("=== END REQUEST DETAILS ===")

	// Send request
	resp, err := s.httpClient.Do(httpReq)
	if err != nil {
		log.Printf("Error forwarding request to %s: %v", fullURL, err)
		return err
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Log complete response details
	log.Printf("=== INCOMING RESPONSE DETAILS ===")
	log.Printf("Status Code: %d", resp.StatusCode)
	log.Printf("Status: %s", resp.Status)
	log.Printf("Response Headers:")
	for name, values := range resp.Header {
		for _, value := range values {
			log.Printf("  %s: %s", name, value)
		}
	}
	log.Printf("Response Body: %s", string(respBody))
	log.Printf("=== END RESPONSE DETAILS ===")

	return nil
}

// ForwardSearchRequest forwards the search request to BPP
func (s *OndcService) ForwardSearchRequest(req interface{}) error {
	return s.forwardRequest("/search", req)
}

// ForwardSelectRequest forwards the select request to BPP
func (s *OndcService) ForwardSelectRequest(req interface{}) error {
	return s.forwardRequest("/select", req)
}

// ForwardInitRequest forwards the init request to BPP
func (s *OndcService) ForwardInitRequest(req interface{}) error {
	return s.forwardRequest("/init", req)
}

// ForwardConfirmRequest forwards the confirm request to BPP
func (s *OndcService) ForwardConfirmRequest(req interface{}) error {
	return s.forwardRequest("/confirm", req)
}

// ForwardStatusRequest forwards the status request to BPP
func (s *OndcService) ForwardStatusRequest(req interface{}) error {
	return s.forwardRequest("/status", req)
}

// ForwardCancelRequest forwards the cancel request to BPP
func (s *OndcService) ForwardCancelRequest(req interface{}) error {
	return s.forwardRequest("/cancel", req)
}

// ForwardUpdateRequest forwards the update request to BPP
func (s *OndcService) ForwardUpdateRequest(req interface{}) error {
	return s.forwardRequest("/update", req)
}

// ForwardTrackRequest forwards the track request to BPP
func (s *OndcService) ForwardTrackRequest(req interface{}) error {
	return s.forwardRequest("/track", req)
}
