package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ondc-buyer-app/internal/config"
)

// SellerService handles seller-side business logic
type SellerService struct {
	config     *config.Config
	httpClient *http.Client
}

// NewSellerService creates a new instance of SellerService
func NewSellerService(cfg *config.Config) *SellerService {
	return &SellerService{
		config:     cfg,
		httpClient: &http.Client{},
	}
}

// sendCallback sends a callback response to the buyer's BAP URI
func (s *SellerService) sendCallback(bapURI string, action string, response interface{}) error {
	jsonData, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("failed to marshal response: %v", err)
	}

	// Construct the callback URL
	callbackURL := fmt.Sprintf("%s/on_%s", bapURI, action)

	// Create and send request
	resp, err := s.httpClient.Post(callbackURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send callback: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read callback response: %v", err)
	}

	return nil
}

// getBapURI extracts the BAP URI from the request context
func (s *SellerService) getBapURI(request map[string]interface{}) (string, error) {
	context, ok := request["context"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid context in request")
	}

	bapURI, ok := context["bap_uri"].(string)
	if !ok {
		return "", fmt.Errorf("bap_uri not found in context")
	}

	return bapURI, nil
}

// ForwardToBapURI forwards the request to the BAP URI found in the request context
func (s *SellerService) ForwardToBapURI(request map[string]interface{}) error {
	context, ok := request["context"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid context in request")
	}

	bapURI, ok := context["bap_uri"].(string)
	if !ok {
		return fmt.Errorf("bap_uri not found in context")
	}

	action, ok := context["action"].(string)
	if !ok {
		return fmt.Errorf("action not found in context")
	}

	// Construct the full callback URL by appending action to bap_uri
	callbackURL := fmt.Sprintf("%s/%s", bapURI, action)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %v", err)
	}

	// Log the callback URL
	fmt.Printf("Forwarding callback to: %s\n", callbackURL)

	// Forward the entire request to the constructed callback URL
	resp, err := s.httpClient.Post(callbackURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to forward request to BAP: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read BAP response: %v", err)
	}

	fmt.Printf("Response from BAP %s: %s\n", callbackURL, string(respBody))
	return nil
}

// ProcessSearchRequest processes the search request and sends on_search callback
func (s *SellerService) ProcessSearchRequest(request map[string]interface{}) error {
	bapURI, err := s.getBapURI(request)
	if err != nil {
		return err
	}

	// TODO: Process search request and prepare catalog response
	response := map[string]interface{}{
		"context": request["context"],
		"message": map[string]interface{}{
			"catalog": map[string]interface{}{
				// Add your catalog items here
			},
		},
	}

	return s.sendCallback(bapURI, "search", response)
}

// ProcessSelectRequest processes the select request and sends on_select callback
func (s *SellerService) ProcessSelectRequest(request map[string]interface{}) error {
	bapURI, err := s.getBapURI(request)
	if err != nil {
		return err
	}

	// TODO: Process select request and prepare order details
	response := map[string]interface{}{
		"context": request["context"],
		"message": map[string]interface{}{
			"order": map[string]interface{}{
				// Add selected order details here
			},
		},
	}

	return s.sendCallback(bapURI, "select", response)
}

// ProcessInitRequest processes the init request and sends on_init callback
func (s *SellerService) ProcessInitRequest(request map[string]interface{}) error {
	bapURI, err := s.getBapURI(request)
	if err != nil {
		return err
	}

	// TODO: Process init request and prepare order details
	response := map[string]interface{}{
		"context": request["context"],
		"message": map[string]interface{}{
			"order": map[string]interface{}{
				// Add initialized order details here
			},
		},
	}

	return s.sendCallback(bapURI, "init", response)
}

// ProcessConfirmRequest processes the confirm request and sends on_confirm callback
func (s *SellerService) ProcessConfirmRequest(request map[string]interface{}) error {
	bapURI, err := s.getBapURI(request)
	if err != nil {
		return err
	}

	// TODO: Process confirm request and prepare order confirmation
	response := map[string]interface{}{
		"context": request["context"],
		"message": map[string]interface{}{
			"order": map[string]interface{}{
				// Add confirmed order details here
			},
		},
	}

	return s.sendCallback(bapURI, "confirm", response)
}

// ProcessStatusRequest processes the status request and sends on_status callback
func (s *SellerService) ProcessStatusRequest(request map[string]interface{}) error {
	bapURI, err := s.getBapURI(request)
	if err != nil {
		return err
	}

	// TODO: Process status request and prepare order status
	response := map[string]interface{}{
		"context": request["context"],
		"message": map[string]interface{}{
			"order": map[string]interface{}{
				// Add order status details here
			},
		},
	}

	return s.sendCallback(bapURI, "status", response)
}

// ProcessCancelRequest processes the cancel request and sends on_cancel callback
func (s *SellerService) ProcessCancelRequest(request map[string]interface{}) error {
	bapURI, err := s.getBapURI(request)
	if err != nil {
		return err
	}

	// TODO: Process cancel request and prepare cancellation confirmation
	response := map[string]interface{}{
		"context": request["context"],
		"message": map[string]interface{}{
			"order": map[string]interface{}{
				// Add cancelled order details here
			},
		},
	}

	return s.sendCallback(bapURI, "cancel", response)
}

// ProcessUpdateRequest processes the update request and sends on_update callback
func (s *SellerService) ProcessUpdateRequest(request map[string]interface{}) error {
	bapURI, err := s.getBapURI(request)
	if err != nil {
		return err
	}

	// TODO: Process update request and prepare update confirmation
	response := map[string]interface{}{
		"context": request["context"],
		"message": map[string]interface{}{
			"order": map[string]interface{}{
				// Add updated order details here
			},
		},
	}

	return s.sendCallback(bapURI, "update", response)
}

// ProcessTrackRequest processes the track request and sends on_track callback
func (s *SellerService) ProcessTrackRequest(request map[string]interface{}) error {
	bapURI, err := s.getBapURI(request)
	if err != nil {
		return err
	}

	// TODO: Process track request and prepare tracking details
	response := map[string]interface{}{
		"context": request["context"],
		"message": map[string]interface{}{
			"tracking": map[string]interface{}{
				// Add tracking details here
			},
		},
	}

	return s.sendCallback(bapURI, "track", response)
}
