package middleware

import (
	"strings"

	"github.com/Alwin18/wms/config"
	"github.com/Alwin18/wms/internal/utils"
	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware validates JWT token
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.SendUnauthorized(c, "Missing authorization header")
		}

		token, err := utils.ExtractTokenFromHeader(authHeader)
		if err != nil {
			return utils.SendUnauthorized(c, "Invalid authorization header format")
		}

		claims, err := utils.ValidateToken(token, config.GlobalConfig.JWT.Secret)
		if err != nil {
			return utils.SendUnauthorized(c, "Invalid or expired token")
		}

		// Store user info in context
		c.Locals("userID", claims.UserID)
		c.Locals("email", claims.Email)
		c.Locals("role", claims.Role)

		return c.Next()
	}
}

// RoleMiddleware checks if user has required role
func RoleMiddleware(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role")
		if role == nil {
			return utils.SendForbidden(c, "Access denied")
		}

		userRole := role.(string)
		for _, allowedRole := range allowedRoles {
			if strings.EqualFold(userRole, allowedRole) {
				return c.Next()
			}
		}

		return utils.SendForbidden(c, "Insufficient permissions")
	}
}
