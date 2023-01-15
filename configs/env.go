package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	// Load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	mongoIRU := os.Getenv("MONGOURI")
	return mongoIRU
}
