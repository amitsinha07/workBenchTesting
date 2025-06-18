package handlers

import (
	"ondc-buyer-app/internal/config"
	"ondc-buyer-app/internal/services"
	"ondc-buyer-app/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// SellerCallbackHandlers contains all seller callback endpoint handlers
type SellerCallbackHandlers struct {
	sellerService *services.SellerService
}

// NewSellerCallbackHandlers creates a new instance of SellerCallbackHandlers
func NewSellerCallbackHandlers(cfg *config.Config) *SellerCallbackHandlers {
	return &SellerCallbackHandlers{
		sellerService: services.NewSellerService(cfg),
	}
}

// createACKResponse creates a standard ONDC ACK response
func (h *SellerCallbackHandlers) createACKResponse() fiber.Map {
	return fiber.Map{
		"message": fiber.Map{
			"ack": fiber.Map{
				"status": "ACK",
				"number": utils.GenerateACKNumber(),
			},
		},
	}
}

// createNACKResponse creates a standard ONDC NACK response with error
func (h *SellerCallbackHandlers) createNACKResponse(errorMsg string) fiber.Map {
	return fiber.Map{
		"message": fiber.Map{
			"ack": fiber.Map{
				"status": "NACK",
				"number": utils.GenerateACKNumber(),
			},
		},
		"error": fiber.Map{
			"type":    "CORE-ERROR",
			"code":    "10001",
			"message": errorMsg,
		},
	}
}

// HandleOnSearch handles the on_search callback
func (h *SellerCallbackHandlers) HandleOnSearch(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse("Invalid request format"))
	}

	if err := h.sellerService.ForwardToBapURI(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(h.createNACKResponse("Failed to forward on_search request"))
	}

	return c.JSON(h.createACKResponse())
}

// HandleOnSelect handles the on_select callback
func (h *SellerCallbackHandlers) HandleOnSelect(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse("Invalid request format"))
	}

	if err := h.sellerService.ForwardToBapURI(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(h.createNACKResponse("Failed to forward on_select request"))
	}

	return c.JSON(h.createACKResponse())
}

// HandleOnInit handles the on_init callback
func (h *SellerCallbackHandlers) HandleOnInit(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse("Invalid request format"))
	}

	if err := h.sellerService.ForwardToBapURI(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(h.createNACKResponse("Failed to forward on_init request"))
	}

	return c.JSON(h.createACKResponse())
}

// HandleOnConfirm handles the on_confirm callback
func (h *SellerCallbackHandlers) HandleOnConfirm(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse("Invalid request format"))
	}

	if err := h.sellerService.ForwardToBapURI(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(h.createNACKResponse("Failed to forward on_confirm request"))
	}

	return c.JSON(h.createACKResponse())
}

// HandleOnStatus handles the on_status callback
func (h *SellerCallbackHandlers) HandleOnStatus(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse("Invalid request format"))
	}

	if err := h.sellerService.ForwardToBapURI(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(h.createNACKResponse("Failed to forward on_status request"))
	}

	return c.JSON(h.createACKResponse())
}

// HandleOnCancel handles the on_cancel callback
func (h *SellerCallbackHandlers) HandleOnCancel(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse("Invalid request format"))
	}

	if err := h.sellerService.ForwardToBapURI(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(h.createNACKResponse("Failed to forward on_cancel request"))
	}

	return c.JSON(h.createACKResponse())
}

// HandleOnUpdate handles the on_update callback
func (h *SellerCallbackHandlers) HandleOnUpdate(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse("Invalid request format"))
	}

	if err := h.sellerService.ForwardToBapURI(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(h.createNACKResponse("Failed to forward on_update request"))
	}

	return c.JSON(h.createACKResponse())
}

// HandleOnTrack handles the on_track callback
func (h *SellerCallbackHandlers) HandleOnTrack(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.createNACKResponse("Invalid request format"))
	}

	if err := h.sellerService.ForwardToBapURI(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(h.createNACKResponse("Failed to forward on_track request"))
	}

	return c.JSON(h.createACKResponse())
}
