package config

import (
	"log"
	"os"
)

func getEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatalf("%s is not set", key)
	}

	return value
}
