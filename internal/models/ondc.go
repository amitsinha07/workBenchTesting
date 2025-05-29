package models

import "time"

// Location represents the location structure in ONDC context
type Location struct {
	Country struct {
		Code string `json:"code"`
	} `json:"country"`
	City struct {
		Code string `json:"code"`
	} `json:"city"`
}

// ONDCContext represents the ONDC context object
type ONDCContext struct {
	Location      Location  `json:"location"`
	Domain        string    `json:"domain"`
	Timestamp     time.Time `json:"timestamp"`
	BapID         string    `json:"bap_id"`
	TransactionID string    `json:"transaction_id"`
	MessageID     string    `json:"message_id"`
	Version       string    `json:"version"`
	Action        string    `json:"action"`
	BapURI        string    `json:"bap_uri"`
	TTL           string    `json:"ttl"`
}

// Descriptor represents a generic descriptor
type Descriptor struct {
	Code string `json:"code"`
}

// TagList represents a list item in a tag
type TagList struct {
	Descriptor Descriptor `json:"descriptor"`
	Value      string     `json:"value"`
}

// Tag represents a tag in the payment structure
type Tag struct {
	Descriptor Descriptor `json:"descriptor"`
	Display    bool       `json:"display"`
	List       []TagList  `json:"list"`
}

// Vehicle represents the vehicle details in fulfillment
type Vehicle struct {
	Category string `json:"category"`
}

// Fulfillment represents the fulfillment details
type Fulfillment struct {
	Vehicle Vehicle `json:"vehicle"`
}

// Payment represents the payment details
type Payment struct {
	CollectedBy string `json:"collected_by"`
	Tags        []Tag  `json:"tags"`
}

// Intent represents the search intent
type Intent struct {
	Fulfillment Fulfillment `json:"fulfillment"`
	Payment     Payment     `json:"payment"`
}

// SearchMessage represents the message in search request
type SearchMessage struct {
	Intent Intent `json:"intent"`
}

// SearchRequest represents the complete search request
type SearchRequest struct {
	Context ONDCContext   `json:"context"`
	Message SearchMessage `json:"message"`
}

// Response represents a standard API response
type Response struct {
	Message string                 `json:"message"`
	Ack     map[string]string      `json:"ack,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Error   string                 `json:"error,omitempty"`
}
