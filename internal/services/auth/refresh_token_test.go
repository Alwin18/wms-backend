package auth

import (
	"testing"
	"time"

	"github.com/Alwin18/wms/config"
	"github.com/Alwin18/wms/internal/dto"
	"github.com/Alwin18/wms/internal/models"
	"github.com/Alwin18/wms/internal/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestRefreshToken(t *testing.T) {
	// Initialize config and DB
	app, err := config.Bootstrap()
	if err != nil {
		t.Skipf("Skipping test - database not available: %v", err)
		return
	}

	// Setup: Create a test user
	password := "password123"
	hashedPassword, _ := utils.HashPassword(password)
	user := models.User{
		Email:    "test_refresh@example.com",
		Password: hashedPassword,
		Name:     "Test Refresh User",
		Status:   "active",
		RoleID:   1,
	}

	// Ensure role exists
	var role models.Role
	if err := app.DB.First(&role, 1).Error; err != nil {
		role = models.Role{Name: "admin", Description: "Admin Role"}
		app.DB.Create(&role)
		user.RoleID = role.ID
	}

	// Clean up first
	app.DB.Unscoped().Where("email = ?", user.Email).Delete(&models.User{})

	// Create user
	if err := app.DB.Create(&user).Error; err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	// Teardown
	defer app.DB.Unscoped().Delete(&user)

	// Helper to generate token
	generateToken := func(userID uint, secret string, exp time.Duration) string {
		token, _ := utils.GenerateRefreshToken(userID, secret, exp)
		return token
	}

	// Test Case 1: Success
	t.Run("Success", func(t *testing.T) {
		token := generateToken(user.ID, config.GlobalConfig.JWT.Secret, time.Hour)
		req := &dto.RefreshTokenRequest{
			RefreshToken: token,
		}

		resp, err := RefreshToken(req)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.AccessToken)
	})

	// Test Case 2: Invalid Token Format
	t.Run("InvalidTokenFormat", func(t *testing.T) {
		req := &dto.RefreshTokenRequest{
			RefreshToken: "invalid-token-format",
		}

		resp, err := RefreshToken(req)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	// Test Case 3: Expired Token
	t.Run("ExpiredToken", func(t *testing.T) {
		// Manually create expired token
		claims := jwt.MapClaims{
			"sub": "123",
			"exp": time.Now().Add(-time.Hour).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString([]byte(config.GlobalConfig.JWT.Secret))

		req := &dto.RefreshTokenRequest{
			RefreshToken: tokenString,
		}

		resp, err := RefreshToken(req)
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "invalid or expired refresh token")
	})

	// Test Case 4: Invalid Signature
	t.Run("InvalidSignature", func(t *testing.T) {
		token := generateToken(user.ID, "wrong-secret", time.Hour)
		req := &dto.RefreshTokenRequest{
			RefreshToken: token,
		}

		resp, err := RefreshToken(req)
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "invalid or expired refresh token")
	})

	// Test Case 5: User Not Found
	t.Run("UserNotFound", func(t *testing.T) {
		// Generate token for non-existent user ID
		token := generateToken(999999, config.GlobalConfig.JWT.Secret, time.Hour)
		req := &dto.RefreshTokenRequest{
			RefreshToken: token,
		}

		resp, err := RefreshToken(req)
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, "user not found", err.Error())
	})

	// Test Case 6: Inactive User
	t.Run("InactiveUser", func(t *testing.T) {
		// Create inactive user
		inactiveUser := models.User{
			Email:    "inactive@example.com",
			Password: hashedPassword,
			Name:     "Inactive User",
			Status:   "inactive",
			RoleID:   1,
		}
		app.DB.Create(&inactiveUser)
		defer app.DB.Unscoped().Delete(&inactiveUser)

		token := generateToken(inactiveUser.ID, config.GlobalConfig.JWT.Secret, time.Hour)
		req := &dto.RefreshTokenRequest{
			RefreshToken: token,
		}

		resp, err := RefreshToken(req)
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, "user account is not active", err.Error())
	})

	// Test Case 7: Invalid Claims (Missing sub)
	t.Run("InvalidClaims", func(t *testing.T) {
		claims := jwt.MapClaims{
			"exp": time.Now().Add(time.Hour).Unix(),
			// Missing "sub"
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString([]byte(config.GlobalConfig.JWT.Secret))

		req := &dto.RefreshTokenRequest{
			RefreshToken: tokenString,
		}

		resp, err := RefreshToken(req)
		assert.Error(t, err)
		assert.Nil(t, resp)
		// Note: The error message might vary depending on where it fails first,
		// but typically it should fail at claim extraction or validation
	})
}
