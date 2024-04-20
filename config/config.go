package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// config function to get environment variables

func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("error loading .env file")
	}
	return os.Getenv(key)
}
