package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Log      LogConfig
	CORS     CORSConfig
}

type AppConfig struct {
	Name  string
	Env   string
	Port  string
	Debug bool
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	Timezone string
}

type JWTConfig struct {
	Secret          string
	AccessTokenExp  time.Duration
	RefreshTokenExp time.Duration
}

type LogConfig struct {
	Level  string
	Format string
}

type CORSConfig struct {
	AllowedOrigins string
}

var GlobalConfig *Config

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	// Load .env file if exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Parse JWT expiration durations
	accessTokenExp, err := time.ParseDuration(getEnv("JWT_ACCESS_TOKEN_EXPIRE", "15m"))
	if err != nil {
		log.Fatal("Invalid JWT_ACCESS_TOKEN_EXPIRE format")
	}

	refreshTokenExp, err := time.ParseDuration(getEnv("JWT_REFRESH_TOKEN_EXPIRE", "168h"))
	if err != nil {
		log.Fatal("Invalid JWT_REFRESH_TOKEN_EXPIRE format")
	}

	config := &Config{
		App: AppConfig{
			Name:  getEnv("APP_NAME", "WMS Backend"),
			Env:   getEnv("APP_ENV", "development"),
			Port:  getEnv("APP_PORT", "8080"),
			Debug: getEnvAsBool("APP_DEBUG", true),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "password"),
			DBName:   getEnv("DB_NAME", "wms_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
			Timezone: getEnv("DB_TIMEZONE", "Asia/Jakarta"),
		},
		JWT: JWTConfig{
			Secret:          getEnv("JWT_SECRET", "your-secret-key"),
			AccessTokenExp:  accessTokenExp,
			RefreshTokenExp: refreshTokenExp,
		},
		Log: LogConfig{
			Level:  getEnv("LOG_LEVEL", "debug"),
			Format: getEnv("LOG_FORMAT", "json"),
		},
		CORS: CORSConfig{
			AllowedOrigins: getEnv("CORS_ALLOWED_ORIGINS", "*"),
		},
	}

	GlobalConfig = config
	return config
}

// Helper functions
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valStr := getEnv(key, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultValue
}
