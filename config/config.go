package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var cfg *Config //Config → type, config / cfg → value

type DBConfig struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	EnableSSLMode bool
}

type Config struct {
	Version      string
	ServiceName  string
	HTTPPort     int
	JWTSecretKey string
	DB           *DBConfig
}

func loadConfig() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("Failed to load the env variable: %w", err)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		return fmt.Errorf("VERSION is required")
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		return fmt.Errorf("SERVICE_NAME is required")
	}

	HTTPPort := os.Getenv("HTTP_PORT")
	if HTTPPort == "" {
		return fmt.Errorf("HTTP_PORT is required")
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		return fmt.Errorf("JWT_SECRET_KEY is required")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return fmt.Errorf("DB_HOST is required")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		return fmt.Errorf("DB_PORT is required")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return fmt.Errorf("DB_NAME is required")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return fmt.Errorf("DB_USER is required")
	}

	dbPass := os.Getenv("DB_PASSWORD")
	if dbPass == "" {
		return fmt.Errorf("DB_PASSWORD is required")
	}

	enableSslMode := os.Getenv("DB_ENABLE_SSL_MODE")
	enableSSL, err := strconv.ParseBool(enableSslMode)
	if err != nil {
		return fmt.Errorf("invalid DB_ENABLE_SSL_MODE: %w", err)
	}

	dbPortInt, err := strconv.Atoi(dbPort)
	if err != nil {
		return fmt.Errorf("invalid DB_PORT: %w", err)
	}

	port, err := strconv.Atoi(HTTPPort)
	if err != nil {
		return fmt.Errorf("invalid HTTP_PORT: %w", err)
	}

	// Assign to cfg only after everything is valid
	cfg = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HTTPPort:     port,
		JWTSecretKey: jwtSecretKey,
		DB: &DBConfig{
			Host:          dbHost,
			Port:          dbPortInt,
			Name:          dbName,
			User:          dbUser,
			Password:      dbPass,
			EnableSSLMode: enableSSL,
		},
	}

	return nil
}

func GetConfig() *Config {
	if cfg == nil { // this will stop repetition of loading config
		loadConfig()
	}
	return cfg
}
