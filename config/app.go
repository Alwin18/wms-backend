package config

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Application struct {
	Config    *Config
	DB        *gorm.DB
	Logger    *logrus.Logger
	Validator *validator.Validate
	Fiber     *fiber.App
}

// Bootstrap initializes all application dependencies
func Bootstrap() (*Application, error) {
	// Load configuration
	cfg := LoadConfig()
	log.Println("Configuration loaded successfully")

	// Initialize logger
	logger := InitLogger(cfg)
	logger.Info("Logger initialized successfully")

	// Initialize database
	db, err := InitDatabase(cfg)
	if err != nil {
		return nil, err
	}

	// Initialize validator
	validate := InitValidator()
	logger.Info("Validator initialized successfully")

	// Initialize Fiber app
	app := InitFiber(cfg)
	logger.Info("Fiber app initialized successfully")

	return &Application{
		Config:    cfg,
		DB:        db,
		Logger:    logger,
		Validator: validate,
		Fiber:     app,
	}, nil
}
