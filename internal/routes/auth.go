package routes

import (
	"github.com/Alwin18/wms/internal/handlers"
	"github.com/Alwin18/wms/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

// SetupAuthRoutes sets up authentication routes
func SetupAuthRoutes(app *fiber.App) {
	authHandler := handlers.NewAuthHandler()

	// Auth routes group
	auth := app.Group("/api/v1/auth")

	// Public routes (no authentication required)
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	auth.Post("/refresh-token", authHandler.RefreshToken)
	auth.Post("/forgot-password", authHandler.ForgotPassword)
	auth.Post("/reset-password", authHandler.ResetPassword)

	// Protected routes (authentication required)
	auth.Post("/logout", middleware.AuthMiddleware(), authHandler.Logout)
	auth.Post("/change-password", middleware.AuthMiddleware(), authHandler.ChangePassword)
}
