package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %v", err)
	}

	dbUsername, ok := os.LookupEnv("DB_USERNAME")
	if !ok {
		return nil, fmt.Errorf("DB_USERNAME environment variable is not set")
	}

	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return nil, fmt.Errorf("DB_PASSWORD environment variable is not set")
	}

	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return nil, fmt.Errorf("DB_HOST environment variable is not set")
	}

	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return nil, fmt.Errorf("DB_PORT environment variable is not set")
	}

	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return nil, fmt.Errorf("DB_NAME environment variable is not set")
	}

	return &Config{
		DBUsername: dbUsername,
		DBPassword: dbPassword,
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBName:     dbName,
	}, nil
}

func (c *Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUsername, c.DBPassword, c.DBHost, c.DBPort, c.DBName,
	)
}
