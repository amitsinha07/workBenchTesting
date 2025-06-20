package routes

import (
	"ondc-buyer-app/internal/api/handlers"
	"ondc-buyer-app/internal/config"
	"ondc-buyer-app/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(app *fiber.App, cfg *config.Config) {
	// Create handlers
	actionHandlers := handlers.NewActionHandlers(cfg)
	callbackHandlers := handlers.NewCallbackHandlers()
	sellerHandlers := handlers.NewSellerHandlers(cfg)
	sellerCallbackHandlers := handlers.NewSellerCallbackHandlers(cfg)

	// Create auth middleware instance
	authMiddleware := middleware.NewAuthMiddleware()

	// API v1 routes
	v1 := app.Group("/v1")

	// Buyer endpoints (action APIs - outgoing requests)
	// These don't need auth verification since they're incoming requests to our API
	// The auth headers are added automatically in the OndcService when forwarding
	buyer := v1.Group("/buyer")
	buyer.Post("/search", actionHandlers.HandleSearch)
	buyer.Post("/select", actionHandlers.HandleSelect)
	buyer.Post("/init", actionHandlers.HandleInit)
	buyer.Post("/confirm", actionHandlers.HandleConfirm)
	buyer.Post("/status", actionHandlers.HandleStatus)
	buyer.Post("/cancel", actionHandlers.HandleCancel)
	buyer.Post("/update", actionHandlers.HandleUpdate)
	buyer.Post("/track", actionHandlers.HandleTrack)

	// Buyer callback endpoints (incoming callbacks from BPP)
	// Use optional auth verification - verify if auth header is present
	buyerCallbacks := buyer.Group("", authMiddleware.OptionalVerifyAuthHeader())
	buyerCallbacks.Post("/on_search", callbackHandlers.HandleOnSearch)
	buyerCallbacks.Post("/on_select", callbackHandlers.HandleOnSelect)
	buyerCallbacks.Post("/on_init", callbackHandlers.HandleOnInit)
	buyerCallbacks.Post("/on_confirm", callbackHandlers.HandleOnConfirm)
	buyerCallbacks.Post("/on_status", callbackHandlers.HandleOnStatus)
	buyerCallbacks.Post("/on_cancel", callbackHandlers.HandleOnCancel)
	buyerCallbacks.Post("/on_update", callbackHandlers.HandleOnUpdate)
	buyerCallbacks.Post("/on_track", callbackHandlers.HandleOnTrack)

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

	// Seller callback endpoints (incoming callbacks)
	// Use optional auth verification - verify if auth header is present
	sellerCallbacks := seller.Group("", authMiddleware.OptionalVerifyAuthHeader())
	sellerCallbacks.Post("/on_search", sellerCallbackHandlers.HandleOnSearch)
	sellerCallbacks.Post("/on_select", sellerCallbackHandlers.HandleOnSelect)
	sellerCallbacks.Post("/on_init", sellerCallbackHandlers.HandleOnInit)
	sellerCallbacks.Post("/on_confirm", sellerCallbackHandlers.HandleOnConfirm)
	sellerCallbacks.Post("/on_status", sellerCallbackHandlers.HandleOnStatus)
	sellerCallbacks.Post("/on_cancel", sellerCallbackHandlers.HandleOnCancel)
	sellerCallbacks.Post("/on_update", sellerCallbackHandlers.HandleOnUpdate)
	sellerCallbacks.Post("/on_track", sellerCallbackHandlers.HandleOnTrack)
}
