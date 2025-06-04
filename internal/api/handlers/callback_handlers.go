package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// CallbackHandlers contains all callback endpoint handlers
type CallbackHandlers struct{}

// NewCallbackHandlers creates a new instance of CallbackHandlers
func NewCallbackHandlers() *CallbackHandlers {
	return &CallbackHandlers{}
}

// HandleOnSearch handles the on_search callback
func (h *CallbackHandlers) HandleOnSearch(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.Map{
				"ack": fiber.Map{
					"status": "NACK",
				},
			},
		})
	}

	log.Printf("Received on_search callback: %+v", request)

	return c.JSON(fiber.Map{
		"message": fiber.Map{
			"ack": fiber.Map{
				"status": "ACK",
			},
		},
	})
}

// HandleOnSelect handles the on_select callback
func (h *CallbackHandlers) HandleOnSelect(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.Map{
				"ack": fiber.Map{
					"status": "NACK",
				},
			},
		})
	}

	log.Printf("Received on_select callback: %+v", request)

	return c.JSON(fiber.Map{
		"message": fiber.Map{
			"ack": fiber.Map{
				"status": "ACK",
			},
		},
	})
}

// HandleOnInit handles the on_init callback
func (h *CallbackHandlers) HandleOnInit(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.Map{
				"ack": fiber.Map{
					"status": "NACK",
				},
			},
		})
	}

	log.Printf("Received on_init callback: %+v", request)

	return c.JSON(fiber.Map{
		"message": fiber.Map{
			"ack": fiber.Map{
				"status": "ACK",
			},
		},
	})
}

// HandleOnConfirm handles the on_confirm callback
func (h *CallbackHandlers) HandleOnConfirm(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.Map{
				"ack": fiber.Map{
					"status": "NACK",
				},
			},
		})
	}

	log.Printf("Received on_confirm callback: %+v", request)

	return c.JSON(fiber.Map{
		"message": fiber.Map{
			"ack": fiber.Map{
				"status": "ACK",
			},
		},
	})
}

// HandleOnStatus handles the on_status callback
func (h *CallbackHandlers) HandleOnStatus(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.Map{
				"ack": fiber.Map{
					"status": "NACK",
				},
			},
		})
	}

	log.Printf("Received on_status callback: %+v", request)

	return c.JSON(fiber.Map{
		"message": fiber.Map{
			"ack": fiber.Map{
				"status": "ACK",
			},
		},
	})
}

// HandleOnCancel handles the on_cancel callback
func (h *CallbackHandlers) HandleOnCancel(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.Map{
				"ack": fiber.Map{
					"status": "NACK",
				},
			},
		})
	}

	log.Printf("Received on_cancel callback: %+v", request)

	return c.JSON(fiber.Map{
		"message": fiber.Map{
			"ack": fiber.Map{
				"status": "ACK",
			},
		},
	})
}

// HandleOnUpdate handles the on_update callback
func (h *CallbackHandlers) HandleOnUpdate(c *fiber.Ctx) error {
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.Map{
				"ack": fiber.Map{
					"status": "NACK",
				},
			},
		})
	}

	log.Printf("Received on_update callback: %+v", request)

	return c.JSON(fiber.Map{
		"message": fiber.Map{
			"ack": fiber.Map{
				"status": "ACK",
			},
		},
	})
}
