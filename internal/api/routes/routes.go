package routes

import (
	"ondc-buyer-app/internal/api/handlers"
	"ondc-buyer-app/internal/config"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(app *fiber.App, cfg *config.Config) {
	// Create handlers
	actionHandlers := handlers.NewActionHandlers(cfg)
	callbackHandlers := handlers.NewCallbackHandlers()

	// API v1 routes
	v1 := app.Group("/v1")

	// Action endpoints
	v1.Post("/search", actionHandlers.HandleSearch)
	v1.Post("/select", actionHandlers.HandleSelect)
	v1.Post("/init", actionHandlers.HandleInit)
	v1.Post("/confirm", actionHandlers.HandleConfirm)
	v1.Post("/status", actionHandlers.HandleStatus)
	v1.Post("/cancel", actionHandlers.HandleCancel)
	v1.Post("/update", actionHandlers.HandleUpdate)

	// Callback endpoints
	v1.Post("/on_search", callbackHandlers.HandleOnSearch)
	v1.Post("/on_select", callbackHandlers.HandleOnSelect)
	v1.Post("/on_init", callbackHandlers.HandleOnInit)
	v1.Post("/on_confirm", callbackHandlers.HandleOnConfirm)
	v1.Post("/on_status", callbackHandlers.HandleOnStatus)
	v1.Post("/on_cancel", callbackHandlers.HandleOnCancel)
	v1.Post("/on_update", callbackHandlers.HandleOnUpdate)
}
