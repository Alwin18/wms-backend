package auth

import (
	"errors"
	"strconv"

	"github.com/Alwin18/wms/config"
	"github.com/Alwin18/wms/internal/dto"
	"github.com/Alwin18/wms/internal/models"
	"github.com/Alwin18/wms/internal/utils"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// RefreshToken generates a new access token from refresh token
func RefreshToken(req *dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {
	// Parse refresh token
	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.GlobalConfig.JWT.Secret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired refresh token")
	}

	// Extract user ID from token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	userIDStr, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.New("invalid user ID in token")
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	// Get user from database
	var user models.User
	if err := config.DB.First(&user, uint(userID)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Check if user is active
	if user.Status != "active" {
		return nil, errors.New("user account is not active")
	}

	// Get user's role
	var role string
	if err := config.DB.Model(&user).Association("Roles").Find(&user.Roles); err == nil && len(user.Roles) > 0 {
		role = user.Roles[0].Name
	} else {
		role = "user"
	}

	// Generate new access token
	accessToken, err := utils.GenerateAccessToken(
		user.ID,
		user.Email,
		role,
		config.GlobalConfig.JWT.Secret,
		config.GlobalConfig.JWT.AccessTokenExp,
	)
	if err != nil {
		return nil, err
	}

	return &dto.RefreshTokenResponse{
		AccessToken: accessToken,
		TokenType:   "Bearer",
		ExpiresIn:   int64(config.GlobalConfig.JWT.AccessTokenExp.Seconds()),
	}, nil
}
