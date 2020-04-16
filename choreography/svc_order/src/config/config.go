package config

import (
	"log"

	"github.com/joho/godotenv"
)

// ConfigEnvironment ...
func ConfigEnvironment(env string) {
	if env == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
