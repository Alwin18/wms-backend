package auth

import (
	"errors"

	"github.com/Alwin18/wms/config"
	"github.com/Alwin18/wms/internal/dto"
	"github.com/Alwin18/wms/internal/models"
	"github.com/Alwin18/wms/internal/utils"
	"gorm.io/gorm"
)

// Login authenticates user and returns tokens
func Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var user models.User

	// Find user by email
	if err := config.DB.Where("email = ?", req.Email).Preload("Role").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	// Check if user is active
	if user.Status != "active" {
		return nil, errors.New("user account is not active")
	}

	// Verify password
	if err := utils.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate access token
	accessToken, err := utils.GenerateAccessToken(
		user.ID,
		user.Email,
		user.Role.Name,
		config.GlobalConfig.JWT.Secret,
		config.GlobalConfig.JWT.AccessTokenExp,
	)
	if err != nil {
		return nil, err
	}

	// Generate refresh token
	refreshToken, err := utils.GenerateRefreshToken(
		user.ID,
		config.GlobalConfig.JWT.Secret,
		config.GlobalConfig.JWT.RefreshTokenExp,
	)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    int64(config.GlobalConfig.JWT.AccessTokenExp.Seconds()),
		User: dto.UserDTO{
			ID:     user.ID,
			Email:  user.Email,
			Name:   user.Name,
			Phone:  user.Phone,
			Status: user.Status,
		},
	}, nil
}
