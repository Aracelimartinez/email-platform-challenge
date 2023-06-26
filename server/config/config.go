package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Credentials struct {
	Username string
	Password string
}

var ZincSearchCredentials Credentials

func init()  {
	err:= loadEnv()
	if err !=nil {
		fmt.Errorf("Failed to initialize de environment variables: %w", err)
	}
}

// Initialize the env variables
func loadEnv() (error) {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		log.Fatalf("Error to load the .env file: %v", err)
		return err
	}

	ZincSearchCredentials = Credentials{
		Username: os.Getenv("ZINCSEARCH_USERNAME"),
		Password: os.Getenv("ZINCSEARCH_PASSWORD"),
	}

	return  nil
}
