package utils

import (
	"github.com/gofiber/fiber/v2"
)

// SuccessResponse represents a successful API response
type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Meta    *Meta  `json:"meta,omitempty"`
}

// ErrorResponse represents an error API response
type ErrorResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

// Meta represents pagination metadata
type Meta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"totalPages"`
}

// SendSuccess sends a success response
func SendSuccess(c *fiber.Ctx, message string, data any) error {
	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// SendSuccessWithMeta sends a success response with pagination metadata
func SendSuccessWithMeta(c *fiber.Ctx, message string, data any, meta *Meta) error {
	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

// SendCreated sends a created response (201)
func SendCreated(c *fiber.Ctx, message string, data any) error {
	return c.Status(fiber.StatusCreated).JSON(SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// SendError sends an error response
func SendError(c *fiber.Ctx, statusCode int, message string, errors ...string) error {
	if len(errors) == 0 {
		errors = []string{message}
	}

	return c.Status(statusCode).JSON(ErrorResponse{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}

// SendBadRequest sends a bad request error (400)
func SendBadRequest(c *fiber.Ctx, message string, errors ...string) error {
	return SendError(c, fiber.StatusBadRequest, message, errors...)
}

// SendUnauthorized sends an unauthorized error (401)
func SendUnauthorized(c *fiber.Ctx, message string) error {
	return SendError(c, fiber.StatusUnauthorized, message)
}

// SendForbidden sends a forbidden error (403)
func SendForbidden(c *fiber.Ctx, message string) error {
	return SendError(c, fiber.StatusForbidden, message)
}

// SendNotFound sends a not found error (404)
func SendNotFound(c *fiber.Ctx, message string) error {
	return SendError(c, fiber.StatusNotFound, message)
}

// SendInternalServerError sends an internal server error (500)
func SendInternalServerError(c *fiber.Ctx, message string) error {
	return SendError(c, fiber.StatusInternalServerError, message)
}

// CalculatePagination calculates pagination metadata
func CalculatePagination(page, limit int, total int64) *Meta {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	totalPages := int(total) / limit
	if int(total)%limit > 0 {
		totalPages++
	}

	return &Meta{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}
}
