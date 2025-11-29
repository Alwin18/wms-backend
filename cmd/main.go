package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Alwin18/wms/config"
	"github.com/Alwin18/wms/internal/models"
	"github.com/Alwin18/wms/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Bootstrap application
	app, err := config.Bootstrap()
	if err != nil {
		log.Fatal("Failed to bootstrap application:", err)
	}

	// Run database migrations
	if err := config.AutoMigrate(
		app.DB,
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.AuditLog{},
		&models.Warehouse{},
	); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Setup routes
	setupRoutes(app.Fiber)

	// Start server
	go func() {
		addr := fmt.Sprintf(":%s", app.Config.App.Port)
		app.Logger.Infof("Server starting on %s", addr)
		if err := app.Fiber.Listen(addr); err != nil {
			app.Logger.Fatal("Failed to start server:", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	app.Logger.Info("Shutting down server...")
	if err := app.Fiber.Shutdown(); err != nil {
		app.Logger.Error("Server forced to shutdown:", err)
	}

	// Close database connection
	sqlDB, err := app.DB.DB()
	if err == nil {
		sqlDB.Close()
	}

	app.Logger.Info("Server exited")
}

// setupRoutes registers all application routes
func setupRoutes(app *fiber.App) {
	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "WMS Backend is running",
		})
	})

	// API routes
	routes.SetupAuthRoutes(app)

	// TODO: Add more routes here
	// routes.SetupUserRoutes(app)
	// routes.SetupProductRoutes(app)
	// routes.SetupWarehouseRoutes(app)
	// etc.

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Route not found",
			"errors":  []string{"The requested endpoint does not exist"},
		})
	})
}
