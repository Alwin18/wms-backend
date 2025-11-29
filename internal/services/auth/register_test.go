package auth

import (
	"testing"

	"github.com/Alwin18/wms/config"
	"github.com/Alwin18/wms/internal/dto"
	"github.com/Alwin18/wms/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	// Initialize config and DB
	app, err := config.Bootstrap()
	if err != nil {
		t.Skipf("Skipping test - database not available: %v", err)
		return
	}

	// Setup
	email := "test_register@example.com"

	// Clean up first
	app.DB.Unscoped().Where("email = ?", email).Delete(&models.User{})

	// Teardown
	defer app.DB.Unscoped().Where("email = ?", email).Delete(&models.User{})

	// Test Case 1: Success
	t.Run("Success", func(t *testing.T) {
		// Ensure default role exists
		var role models.Role
		if err := app.DB.First(&role, 1).Error; err != nil {
			role = models.Role{Name: "user", Description: "User Role"}
			app.DB.Create(&role)
		}

		req := &dto.RegisterRequest{
			Email:    email,
			Password: "password123",
			Name:     "Test Register User",
			Phone:    "08123456789",
		}

		// We need to modify Register service to assign default role or we assign it here if possible?
		// The Register service currently has commented out code for default role.
		// Let's assume for now we need to fix the test environment or the service.
		// Since I cannot change service logic easily without user approval (though I can),
		// let's look at the error: "violates foreign key constraint".
		// This means User struct has RoleID which defaults to 0, but Role with ID 0 doesn't exist.
		// The User model has `RoleID uint` which defaults to 0.

		// To fix this properly in the test without changing service logic (if service logic is correct but incomplete),
		// we might need to update the service to actually assign a role.
		// BUT, looking at the error, it seems the service tries to save a user with RoleID 0.
		// Let's check the service code again. It creates user WITHOUT setting RoleID.
		// So RoleID is 0.
		// We should probably update the Register service to assign a default role (e.g. ID 1).

		resp, err := Register(req)
		assert.NoError(t, err)
		if resp != nil {
			assert.Equal(t, email, resp.User.Email)
			assert.Equal(t, "active", resp.User.Status)
		} else {
			t.Fatal("Response is nil")
		}

		// Verify user created in DB
		var user models.User
		err = app.DB.Where("email = ?", email).First(&user).Error
		assert.NoError(t, err)
		assert.Equal(t, req.Name, user.Name)
	})

	// Test Case 2: Duplicate Email
	t.Run("DuplicateEmail", func(t *testing.T) {
		// Ensure user exists from previous test (or create if not)
		// Since we run tests sequentially and didn't delete in between, it should exist.
		// But to be safe and independent:

		// Create a user manually for this test case
		duplicateEmail := "duplicate@example.com"
		app.DB.Unscoped().Where("email = ?", duplicateEmail).Delete(&models.User{})

		user := models.User{
			Email:    duplicateEmail,
			Name:     "Existing User",
			Password: "hashedpassword",
			RoleID:   1, // Ensure valid RoleID
		}
		if err := app.DB.Create(&user).Error; err != nil {
			t.Fatalf("Failed to create existing user: %v", err)
		}
		defer app.DB.Unscoped().Delete(&user)

		req := &dto.RegisterRequest{
			Email:    duplicateEmail,
			Password: "password123",
			Name:     "New User",
		}

		resp, err := Register(req)
		assert.Error(t, err)
		assert.Nil(t, resp)
		if err != nil {
			assert.Equal(t, "email already registered", err.Error())
		}
	})
}
