package main

import (
	"log"
	"os"

	"github.com/christimiller/gaikeep/cmd"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	// Example of accessing an environment variable
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is not set. Please check your .env file.")
	}

	// Continue with executing the root command
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error executing root command: %s", err)
	}
}
