package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Loaded .env file")

	return os.Getenv("MONGO_URI")
}
