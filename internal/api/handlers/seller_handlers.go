package handlers

import (
	"ondc-buyer-app/internal/config"
	"ondc-buyer-app/internal/services"

	"github.com/gofiber/fiber/v2"
)

// SellerHandlers contains all seller endpoint handlers
type SellerHandlers struct {
	sellerService *services.SellerService
}

// NewSellerHandlers creates a new instance of SellerHandlers
func NewSellerHandlers(cfg *config.Config) *SellerHandlers {
	return &SellerHandlers{
		sellerService: services.NewSellerService(cfg),
	}
}

// createACKResponse creates a standard ONDC ACK response
func (h *SellerHandlers) createACKResponse() fiber.Map {
	return fiber.Map{
		"message": fiber.Map{
			"ack": fiber.Map{
				"status": "ACK",
			},
		},
	}
}

// createNACKResponse creates a standard ONDC NACK response with error
func (h *SellerHandlers) createNACKResponse(errorMsg string) fiber.Map {
	return fiber.Map{
		"message": fiber.Map{
			"ack": fiber.Map{
				"status": "NACK",
			},
		},
		"error": fiber.Map{
			"type":    "CORE-ERROR",
			"code":    "10001",
			"message": errorMsg,
		},
	}
}

// HandleSearch handles the search request from buyer
func (h *SellerHandlers) HandleSearch(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse("Invalid request format"))
	}

	if err := h.sellerService.ProcessSearchRequest(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(h.createNACKResponse("Failed to process search request"))
	}

	return c.JSON(h.createACKResponse())
}

// HandleSelect handles the select request from buyer
func (h *SellerHandlers) HandleSelect(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse("Invalid request format"))
	}

	if err := h.sellerService.ProcessSelectRequest(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(h.createNACKResponse("Failed to process select request"))
	}

	return c.JSON(h.createACKResponse())
}

// HandleInit handles the init request from buyer
func (h *SellerHandlers) HandleInit(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse("Invalid request format"))
	}

	if err := h.sellerService.ProcessInitRequest(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(h.createNACKResponse("Failed to process init request"))
	}

	return c.JSON(h.createACKResponse())
}

// HandleConfirm handles the confirm request from buyer
func (h *SellerHandlers) HandleConfirm(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse("Invalid request format"))
	}

	if err := h.sellerService.ProcessConfirmRequest(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(h.createNACKResponse("Failed to process confirm request"))
	}

	return c.JSON(h.createACKResponse())
}

// HandleStatus handles the status request from buyer
func (h *SellerHandlers) HandleStatus(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse("Invalid request format"))
	}

	if err := h.sellerService.ProcessStatusRequest(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(h.createNACKResponse("Failed to process status request"))
	}

	return c.JSON(h.createACKResponse())
}

// HandleCancel handles the cancel request from buyer
func (h *SellerHandlers) HandleCancel(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse("Invalid request format"))
	}

	if err := h.sellerService.ProcessCancelRequest(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(h.createNACKResponse("Failed to process cancel request"))
	}

	return c.JSON(h.createACKResponse())
}

// HandleUpdate handles the update request from buyer
func (h *SellerHandlers) HandleUpdate(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse("Invalid request format"))
	}

	if err := h.sellerService.ProcessUpdateRequest(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(h.createNACKResponse("Failed to process update request"))
	}

	return c.JSON(h.createACKResponse())
}

// HandleTrack handles the track request from buyer
func (h *SellerHandlers) HandleTrack(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse("Invalid request format"))
	}

	if err := h.sellerService.ProcessTrackRequest(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(h.createNACKResponse("Failed to process track request"))
	}

	return c.JSON(h.createACKResponse())
}
