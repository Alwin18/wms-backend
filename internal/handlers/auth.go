package handlers

import (
	"github.com/Alwin18/wms/config"
	"github.com/Alwin18/wms/internal/dto"
	"github.com/Alwin18/wms/internal/services/auth"
	"github.com/Alwin18/wms/internal/utils"
	"github.com/gofiber/fiber/v2"
)

// AuthHandler handles authentication related requests
type AuthHandler struct{}

// NewAuthHandler creates a new auth handler
func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// Login handles user login
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return utils.SendBadRequest(c, "Invalid request body")
	}

	// Validate request
	if err := config.Validator.Struct(&req); err != nil {
		return utils.SendBadRequest(c, "Validation failed", err.Error())
	}

	// Call service
	resp, err := auth.Login(&req)
	if err != nil {
		return utils.SendUnauthorized(c, err.Error())
	}

	return utils.SendSuccess(c, "Login successful", resp)
}

// Register handles user registration
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return utils.SendBadRequest(c, "Invalid request body")
	}

	// Validate request
	if err := config.Validator.Struct(&req); err != nil {
		return utils.SendBadRequest(c, "Validation failed", err.Error())
	}

	// Call service
	resp, err := auth.Register(&req)
	if err != nil {
		return utils.SendBadRequest(c, err.Error())
	}

	return utils.SendCreated(c, "Registration successful", resp)
}

// RefreshToken handles token refresh
func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	var req dto.RefreshTokenRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return utils.SendBadRequest(c, "Invalid request body")
	}

	// Validate request
	if err := config.Validator.Struct(&req); err != nil {
		return utils.SendBadRequest(c, "Validation failed", err.Error())
	}

	// Call service
	resp, err := auth.RefreshToken(&req)
	if err != nil {
		return utils.SendUnauthorized(c, err.Error())
	}

	return utils.SendSuccess(c, "Token refreshed successfully", resp)
}

// ChangePassword handles password change
func (h *AuthHandler) ChangePassword(c *fiber.Ctx) error {
	var req dto.ChangePasswordRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return utils.SendBadRequest(c, "Invalid request body")
	}

	// Validate request
	if err := config.Validator.Struct(&req); err != nil {
		return utils.SendBadRequest(c, "Validation failed", err.Error())
	}

	// Get user ID from context (set by auth middleware)
	userID := c.Locals("userID").(uint)

	// Call service
	if err := auth.ChangePassword(userID, &req); err != nil {
		return utils.SendBadRequest(c, err.Error())
	}

	return utils.SendSuccess(c, "Password changed successfully", nil)
}

// ForgotPassword handles forgot password request
func (h *AuthHandler) ForgotPassword(c *fiber.Ctx) error {
	var req dto.ForgotPasswordRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return utils.SendBadRequest(c, "Invalid request body")
	}

	// Validate request
	if err := config.Validator.Struct(&req); err != nil {
		return utils.SendBadRequest(c, "Validation failed", err.Error())
	}

	// Call service
	if err := auth.ForgotPassword(&req); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}

	return utils.SendSuccess(c, "Password reset instructions sent to your email", nil)
}

// ResetPassword handles password reset
func (h *AuthHandler) ResetPassword(c *fiber.Ctx) error {
	var req dto.ResetPasswordRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return utils.SendBadRequest(c, "Invalid request body")
	}

	// Validate request
	if err := config.Validator.Struct(&req); err != nil {
		return utils.SendBadRequest(c, "Validation failed", err.Error())
	}

	// Call service
	if err := auth.ResetPassword(&req); err != nil {
		return utils.SendBadRequest(c, err.Error())
	}

	return utils.SendSuccess(c, "Password reset successful", nil)
}

// Logout handles user logout
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	// TODO: Implement token blacklisting if needed
	// For now, client-side token removal is sufficient
	return utils.SendSuccess(c, "Logout successful", nil)
}
