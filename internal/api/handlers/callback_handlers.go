package handlers

import (
	"log"
	"ondc-buyer-app/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// CallbackHandlers contains all callback endpoint handlers
type CallbackHandlers struct{}

// NewCallbackHandlers creates a new instance of CallbackHandlers
func NewCallbackHandlers() *CallbackHandlers {
	return &CallbackHandlers{}
}

// createACKResponse creates a standard ONDC ACK response
func (h *CallbackHandlers) createACKResponse() fiber.Map {
	return fiber.Map{
		"message": fiber.Map{
			"ack": fiber.Map{
				"status": "ACK",
				"number": utils.GenerateACKNumber(),
			},
		},
	}
}

// createNACKResponse creates a standard ONDC NACK response
func (h *CallbackHandlers) createNACKResponse() fiber.Map {
	return fiber.Map{
		"message": fiber.Map{
			"ack": fiber.Map{
				"status": "NACK",
				"number": utils.GenerateACKNumber(),
			},
		},
	}
}

// HandleOnSearch handles the on_search callback
func (h *CallbackHandlers) HandleOnSearch(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse())
	}

	log.Printf("Received on_search callback: %+v", request)

	return c.JSON(h.createACKResponse())
}

// HandleOnSelect handles the on_select callback
func (h *CallbackHandlers) HandleOnSelect(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse())
	}

	log.Printf("Received on_select callback: %+v", request)

	return c.JSON(h.createACKResponse())
}

// HandleOnInit handles the on_init callback
func (h *CallbackHandlers) HandleOnInit(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse())
	}

	log.Printf("Received on_init callback: %+v", request)

	return c.JSON(h.createACKResponse())
}

// HandleOnConfirm handles the on_confirm callback
func (h *CallbackHandlers) HandleOnConfirm(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse())
	}

	log.Printf("Received on_confirm callback: %+v", request)

	return c.JSON(h.createACKResponse())
}

// HandleOnStatus handles the on_status callback
func (h *CallbackHandlers) HandleOnStatus(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse())
	}

	log.Printf("Received on_status callback: %+v", request)

	return c.JSON(h.createACKResponse())
}

// HandleOnCancel handles the on_cancel callback
func (h *CallbackHandlers) HandleOnCancel(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse())
	}

	log.Printf("Received on_cancel callback: %+v", request)

	return c.JSON(h.createACKResponse())
}

// HandleOnUpdate handles the on_update callback
func (h *CallbackHandlers) HandleOnUpdate(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse())
	}

	log.Printf("Received on_update callback: %+v", request)

	return c.JSON(h.createACKResponse())
}

// HandleOnTrack handles the on_track callback
func (h *CallbackHandlers) HandleOnTrack(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse())
	}

	log.Printf("Received on_track callback: %+v", request)

	return c.JSON(h.createACKResponse())
}
