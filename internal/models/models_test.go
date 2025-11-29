package models

import (
	"testing"
	"time"

	"github.com/Alwin18/wms/config"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUserModels(t *testing.T) {
	// Initialize config and DB
	app, err := config.Bootstrap()
	if err != nil {
		t.Skipf("Skipping test - database not available: %v", err)
		return
	}

	// Setup: Ensure a role exists
	role := Role{Name: "test_role", Description: "Test Role"}
	if err := app.DB.Where("name = ?", role.Name).FirstOrCreate(&role).Error; err != nil {
		t.Fatalf("Failed to ensure test role exists: %v", err)
	}
	// We don't delete the role at the end because other tests might need it or it might be complex to clean up if used by other tests running in parallel (though here they are sequential).
	// But to be clean, we can try to delete it at the end of the test function if we created it.
	// For now, let's just leave it or delete it in defer.
	defer app.DB.Unscoped().Delete(&role)

	// Test Case 1: Create Valid User
	t.Run("CreateValidUser", func(t *testing.T) {
		email := "model_test@example.com"
		// Clean up
		app.DB.Unscoped().Where("email = ?", email).Delete(&User{})

		user := User{
			Email:    email,
			Password: "hashedpassword",
			Name:     "Model Test User",
			RoleID:   role.ID,
			Status:   "active",
		}

		err := app.DB.Create(&user).Error
		assert.NoError(t, err)
		assert.NotZero(t, user.ID)

		// Cleanup
		app.DB.Unscoped().Delete(&user)
	})

	// Test Case 2: Unique Email Constraint
	t.Run("UniqueEmail", func(t *testing.T) {
		email := "unique_test@example.com"
		// Clean up
		app.DB.Unscoped().Where("email = ?", email).Delete(&User{})

		user1 := User{
			Email:    email,
			Password: "password",
			Name:     "User 1",
			RoleID:   role.ID,
		}
		app.DB.Create(&user1)
		defer app.DB.Unscoped().Delete(&user1)

		user2 := User{
			Email:    email,
			Password: "password",
			Name:     "User 2",
			RoleID:   role.ID,
		}

		err := app.DB.Create(&user2).Error
		assert.Error(t, err)
		if err != nil {
			assert.Contains(t, err.Error(), "violates unique constraint")
		}
	})

	// Test Case 3: Invalid Role ID (Foreign Key)
	t.Run("InvalidRoleID", func(t *testing.T) {
		user := User{
			Email:    "invalid_role@example.com",
			Password: "password",
			Name:     "Invalid Role User",
			RoleID:   999999, // Non-existent role
		}

		err := app.DB.Create(&user).Error
		assert.Error(t, err)
		if err != nil {
			assert.Contains(t, err.Error(), "violates foreign key constraint")
		}
	})
}

func TestRoleModels(t *testing.T) {
	// Initialize config and DB
	app, err := config.Bootstrap()
	if err != nil {
		t.Skipf("Skipping test - database not available: %v", err)
		return
	}

	t.Run("CreateRole", func(t *testing.T) {
		roleName := "new_test_role"
		app.DB.Unscoped().Where("name = ?", roleName).Delete(&Role{})

		role := Role{
			Name:        roleName,
			Description: "Description",
		}

		err := app.DB.Create(&role).Error
		assert.NoError(t, err)
		assert.NotZero(t, role.ID)

		// Cleanup
		app.DB.Unscoped().Delete(&role)
	})

	t.Run("UniqueRoleName", func(t *testing.T) {
		roleName := "unique_role"
		app.DB.Unscoped().Where("name = ?", roleName).Delete(&Role{})

		role1 := Role{Name: roleName}
		app.DB.Create(&role1)
		defer app.DB.Unscoped().Delete(&role1)

		role2 := Role{Name: roleName}
		err := app.DB.Create(&role2).Error
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "violates unique constraint")
	})
}

func TestAuditLog(t *testing.T) {
	// Initialize config and DB
	app, err := config.Bootstrap()
	if err != nil {
		t.Skipf("Skipping test - database not available: %v", err)
		return
	}

	t.Run("CreateAuditLog", func(t *testing.T) {
		log := AuditLog{
			UserID:     48,
			Action:     "",
			Resource:   "",
			ResourceID: "",
			IPAddress:  "",
			UserAgent:  "",
			Details:    "",
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
			DeletedAt:  gorm.DeletedAt{},
			User:       User{},
		}

		err := app.DB.Create(&log).Error
		assert.NoError(t, err)
		assert.NotZero(t, log.ID)

		// Cleanup
		app.DB.Unscoped().Delete(&log)
	})
}
