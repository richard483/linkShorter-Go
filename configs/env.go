package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load("ENV")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGODB_URI")
}

func EnvMongoDB() string {
	err := godotenv.Load("ENV")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DATABASE")
}
