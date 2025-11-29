package auth

import (
	"errors"

	"github.com/Alwin18/wms/config"
	"github.com/Alwin18/wms/internal/dto"
	"github.com/Alwin18/wms/internal/models"
	"github.com/Alwin18/wms/internal/utils"
	"gorm.io/gorm"
)

// Register creates a new user account
func Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	// Check if email already exists
	var existingUser models.User
	if err := config.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("email already registered")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create user
	user := models.User{
		Email:    req.Email,
		Password: hashedPassword,
		Name:     req.Name,
		Phone:    req.Phone,
		Status:   "active",
		RoleID:   1, // Default role ID (user)
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	// Fetch the role to include in response if needed, or just rely on ID
	// For now, we just proceed.

	return &dto.RegisterResponse{
		User: dto.UserDTO{
			ID:     user.ID,
			Email:  user.Email,
			Name:   user.Name,
			Phone:  user.Phone,
			Status: user.Status,
		},
	}, nil
}
