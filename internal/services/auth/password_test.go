package auth

import (
	"testing"

	"github.com/Alwin18/wms/config"
	"github.com/Alwin18/wms/internal/dto"
	"github.com/Alwin18/wms/internal/models"
	"github.com/Alwin18/wms/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestChangePassword(t *testing.T) {
	// Initialize config and DB
	app, err := config.Bootstrap()
	if err != nil {
		t.Skipf("Skipping test - database not available: %v", err)
		return
	}

	// Setup: Create a test user
	oldPassword := "oldpassword123"
	newPassword := "newpassword123"
	hashedPassword, _ := utils.HashPassword(oldPassword)

	user := models.User{
		Email:    "test_password@example.com",
		Password: hashedPassword,
		Name:     "Test Password User",
		Status:   "active",
		RoleID:   1, // Assuming role ID 1 exists (usually seeded or created)
	}

	// Ensure role exists (if not seeded)
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

	// Test Case 1: Success
	t.Run("Success", func(t *testing.T) {
		req := &dto.ChangePasswordRequest{
			OldPassword: oldPassword,
			NewPassword: newPassword,
		}

		err := ChangePassword(user.ID, req)
		assert.NoError(t, err)

		// Verify password changed in DB
		var updatedUser models.User
		app.DB.First(&updatedUser, user.ID)
		err = utils.VerifyPassword(updatedUser.Password, newPassword)
		assert.NoError(t, err, "New password should be verifiable")
	})

	// Test Case 2: Incorrect Old Password
	t.Run("IncorrectOldPassword", func(t *testing.T) {
		// Reset password first (since it was changed in previous test)
		// But wait, we changed it to newPassword. So let's try to change it again using WRONG old password

		req := &dto.ChangePasswordRequest{
			OldPassword: "wrongpassword",
			NewPassword: "anotherpassword",
		}

		err := ChangePassword(user.ID, req)
		assert.Error(t, err)
		assert.Equal(t, "old password is incorrect", err.Error())
	})

	// Test Case 3: User Not Found
	t.Run("UserNotFound", func(t *testing.T) {
		req := &dto.ChangePasswordRequest{
			OldPassword: oldPassword,
			NewPassword: newPassword,
		}

		err := ChangePassword(999999, req) // Non-existent ID
		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})
}

func TestForgotPassword(t *testing.T) {
	// Initialize config and DB
	_, err := config.Bootstrap()
	if err != nil {
		t.Skipf("Skipping test - database not available: %v", err)
		return
	}

	// Test Case 1: Any email (should return nil as it's a placeholder)
	t.Run("AnyEmail", func(t *testing.T) {
		req := &dto.ForgotPasswordRequest{
			Email: "any@example.com",
		}

		err := ForgotPassword(req)
		assert.NoError(t, err)
	})
}

func TestResetPassword(t *testing.T) {
	// Initialize config and DB
	_, err := config.Bootstrap()
	if err != nil {
		t.Skipf("Skipping test - database not available: %v", err)
		return
	}

	newPassword := "newpassword123"
	hashedPassword, _ := utils.HashPassword(newPassword)

	// Test Case 1: Not Implemented
	t.Run("NotImplemented", func(t *testing.T) {
		req := &dto.ResetPasswordRequest{
			Token:       "any@example.com",
			NewPassword: hashedPassword,
		}

		err := ResetPassword(req)
		assert.Error(t, err)
		assert.Equal(t, "password reset not implemented yet", err.Error())
	})
}
