package configs

import (
	"os"
)

func EnvMongoURI() string {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	return os.Getenv("MONGODB_URI")
}

func EnvMongoDB() string {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	return os.Getenv("DATABASE")
}
