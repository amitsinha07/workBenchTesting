package server

import (
	"ondc-buyer-app/internal/api/routes"
	"ondc-buyer-app/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Server represents the HTTP server
type Server struct {
	app    *fiber.App
	config *config.Config
}

// NewServer creates a new HTTP server instance
func NewServer(cfg *config.Config) *Server {
	app := fiber.New(fiber.Config{
		AppName: "ONDC Buyer App",
	})

	// Add middleware
	app.Use(logger.New())

	return &Server{
		app:    app,
		config: cfg,
	}
}

// SetupRoutes configures all the routes for the application
func (s *Server) SetupRoutes() {
	routes.SetupRoutes(s.app, s.config)
}

// Start starts the HTTP server
func (s *Server) Start(port string) error {
	return s.app.Listen(":" + port)
}
