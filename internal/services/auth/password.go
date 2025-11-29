package auth

import (
	"errors"

	"github.com/Alwin18/wms/config"
	"github.com/Alwin18/wms/internal/dto"
	"github.com/Alwin18/wms/internal/models"
	"github.com/Alwin18/wms/internal/utils"
	"gorm.io/gorm"
)

// ChangePassword changes user's password
func ChangePassword(userID uint, req *dto.ChangePasswordRequest) error {
	var user models.User

	// Get user
	if err := config.DB.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	// Verify old password
	if err := utils.VerifyPassword(user.Password, req.OldPassword); err != nil {
		return errors.New("old password is incorrect")
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	// Update password
	if err := config.DB.Model(&user).Update("password", hashedPassword).Error; err != nil {
		return err
	}

	return nil
}

// ForgotPassword initiates password reset process
func ForgotPassword(req *dto.ForgotPasswordRequest) error {
	var user models.User

	// Check if user exists
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Don't reveal if email exists or not for security
			return nil
		}
		return err
	}

	// TODO: Generate password reset token and send email
	// For now, this is a placeholder
	// You would typically:
	// 1. Generate a unique token
	// 2. Store it in database with expiration
	// 3. Send email with reset link containing the token

	config.Logger.Info("Password reset requested for email: ", req.Email)

	return nil
}

// ResetPassword resets user's password using reset token
func ResetPassword(req *dto.ResetPasswordRequest) error {
	// TODO: Implement password reset logic
	// You would typically:
	// 1. Validate the reset token
	// 2. Check if token is not expired
	// 3. Find user associated with token
	// 4. Update password
	// 5. Invalidate the token

	// Placeholder implementation
	return errors.New("password reset not implemented yet")
}
