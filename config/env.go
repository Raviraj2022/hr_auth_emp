package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName              string
	Port                 string
	DatabaseURL          string
	JWTSecret            string
	AccessTokenExpiry    string
	RefreshTokenExpiry   string
}

var App Config

func LoadEnv() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	App = Config{
		AppName:            getEnv("APP_NAME"),
		Port:               getEnv("PORT"),
		DatabaseURL:        getEnv("DATABASE_URL"),
		JWTSecret:          getEnv("JWT_SECRET"),
		AccessTokenExpiry:  getEnv("ACCESS_TOKEN_EXPIRY"),
		RefreshTokenExpiry: getEnv("REFRESH_TOKEN_EXPIRY"),
	}
}