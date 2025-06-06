package services

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"ondc-buyer-app/internal/config"
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

// forwardRequest is a helper function to forward requests to BPP
func (s *OndcService) forwardRequest(endpoint string, req interface{}) error {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	// Create and send request
	resp, err := s.httpClient.Post(s.config.BppURI+endpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("Response from BPP for %s: %s", endpoint, string(respBody))
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
