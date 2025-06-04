package handlers

import (
	"ondc-buyer-app/internal/config"
	"ondc-buyer-app/internal/services"

	"github.com/gofiber/fiber/v2"
)

// ActionHandlers contains all action endpoint handlers
type ActionHandlers struct {
	ondcService *services.OndcService
}

// NewActionHandlers creates a new instance of ActionHandlers
func NewActionHandlers(cfg *config.Config) *ActionHandlers {
	return &ActionHandlers{
		ondcService: services.NewOndcService(cfg),
	}
}

// HandleSearch handles the search request
func (h *ActionHandlers) HandleSearch(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	if err := h.ondcService.ForwardSearchRequest(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to forward search request",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Search request forwarded successfully",
	})
}

// HandleSelect handles the select request
func (h *ActionHandlers) HandleSelect(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	if err := h.ondcService.ForwardSelectRequest(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to forward select request",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Select request forwarded successfully",
	})
}

// HandleInit handles the init request
func (h *ActionHandlers) HandleInit(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	if err := h.ondcService.ForwardInitRequest(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to forward init request",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Init request forwarded successfully",
	})
}

// HandleConfirm handles the confirm request
func (h *ActionHandlers) HandleConfirm(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	if err := h.ondcService.ForwardConfirmRequest(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to forward confirm request",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Confirm request forwarded successfully",
	})
}

// HandleStatus handles the status request
func (h *ActionHandlers) HandleStatus(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	if err := h.ondcService.ForwardStatusRequest(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to forward status request",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Status request forwarded successfully",
	})
}

// HandleCancel handles the cancel request
func (h *ActionHandlers) HandleCancel(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	if err := h.ondcService.ForwardCancelRequest(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to forward cancel request",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Cancel request forwarded successfully",
	})
}

// HandleUpdate handles the update request
func (h *ActionHandlers) HandleUpdate(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	if err := h.ondcService.ForwardUpdateRequest(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to forward update request",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Update request forwarded successfully",
	})
}
