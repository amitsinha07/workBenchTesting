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
	sellerHandlers := handlers.NewSellerHandlers(cfg)
	sellerCallbackHandlers := handlers.NewSellerCallbackHandlers(cfg)

	// API v1 routes
	v1 := app.Group("/v1")

	// Buyer endpoints
	buyer := v1.Group("/buyer")
	buyer.Post("/search", actionHandlers.HandleSearch)
	buyer.Post("/select", actionHandlers.HandleSelect)
	buyer.Post("/init", actionHandlers.HandleInit)
	buyer.Post("/confirm", actionHandlers.HandleConfirm)
	buyer.Post("/status", actionHandlers.HandleStatus)
	buyer.Post("/cancel", actionHandlers.HandleCancel)
	buyer.Post("/update", actionHandlers.HandleUpdate)
	buyer.Post("/track", actionHandlers.HandleTrack)

	// Buyer callback endpoints
	buyer.Post("/on_search", callbackHandlers.HandleOnSearch)
	buyer.Post("/on_select", callbackHandlers.HandleOnSelect)
	buyer.Post("/on_init", callbackHandlers.HandleOnInit)
	buyer.Post("/on_confirm", callbackHandlers.HandleOnConfirm)
	buyer.Post("/on_status", callbackHandlers.HandleOnStatus)
	buyer.Post("/on_cancel", callbackHandlers.HandleOnCancel)
	buyer.Post("/on_update", callbackHandlers.HandleOnUpdate)
	buyer.Post("/on_track", callbackHandlers.HandleOnTrack)

	// Seller endpoints (mirror of buyer endpoints)
	seller := v1.Group("/seller")
	seller.Post("/search", sellerHandlers.HandleSearch)
	seller.Post("/select", sellerHandlers.HandleSelect)
	seller.Post("/init", sellerHandlers.HandleInit)
	seller.Post("/confirm", sellerHandlers.HandleConfirm)
	seller.Post("/status", sellerHandlers.HandleStatus)
	seller.Post("/cancel", sellerHandlers.HandleCancel)
	seller.Post("/update", sellerHandlers.HandleUpdate)
	seller.Post("/track", sellerHandlers.HandleTrack)

	// Seller callback endpoints
	seller.Post("/on_search", sellerCallbackHandlers.HandleOnSearch)
	seller.Post("/on_select", sellerCallbackHandlers.HandleOnSelect)
	seller.Post("/on_init", sellerCallbackHandlers.HandleOnInit)
	seller.Post("/on_confirm", sellerCallbackHandlers.HandleOnConfirm)
	seller.Post("/on_status", sellerCallbackHandlers.HandleOnStatus)
	seller.Post("/on_cancel", sellerCallbackHandlers.HandleOnCancel)
	seller.Post("/on_update", sellerCallbackHandlers.HandleOnUpdate)
	seller.Post("/on_track", sellerCallbackHandlers.HandleOnTrack)
}
