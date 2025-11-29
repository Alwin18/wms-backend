package auth

import (
	"testing"

	"github.com/Alwin18/wms/config"
	"github.com/Alwin18/wms/internal/dto"
	"github.com/Alwin18/wms/internal/models"
	"github.com/Alwin18/wms/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
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
		Email:    "test_login@example.com",
		Password: hashedPassword,
		Name:     "Test Login User",
		Status:   "active",
		RoleID:   1,
	}

	// Clean up first in case it exists from previous failed runs
	app.DB.Unscoped().Where("email = ?", user.Email).Delete(&models.User{})

	// Create user
	if err := app.DB.Create(&user).Error; err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	// Teardown
	defer app.DB.Unscoped().Delete(&user)

	// Test Case 1: Success
	t.Run("Success", func(t *testing.T) {
		req := &dto.LoginRequest{
			Email:    user.Email,
			Password: password,
		}

		resp, err := Login(req)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		if resp != nil {
			assert.Equal(t, user.Email, resp.User.Email)
			assert.NotEmpty(t, resp.AccessToken)
		}
	})

	// Test Case 2: Wrong Password
	t.Run("WrongPassword", func(t *testing.T) {
		req := &dto.LoginRequest{
			Email:    user.Email,
			Password: "wrongpassword",
		}

		resp, err := Login(req)
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, "invalid email or password", err.Error())
	})

	// Test Case 3: Non-existent User
	t.Run("NonExistentUser", func(t *testing.T) {
		req := &dto.LoginRequest{
			Email:    "nonexistent@example.com",
			Password: "password123",
		}

		resp, err := Login(req)
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, "invalid email or password", err.Error())
	})
}
