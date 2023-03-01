package config

import (
	"log"

	"github.com/joho/godotenv"
)

func InitEnvironment() {
	// loads into ENV
	// to retrieve a environment variable, use os.Getenv("ENV_VARIABLE_NAME")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
