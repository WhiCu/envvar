package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("./config/.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// Get the GITHUB_USERNAME environment variable
	githubUsername, exists := os.LookupEnv("GITHUB_USERNAME")

	if exists {
		fmt.Println(githubUsername)
	}

	// Get the GITHUB_API_KEY environment variable
	githubAPIKey, exists := os.LookupEnv("GITHUB_API_KEY")

	if exists {
		fmt.Println(githubAPIKey)
	}
}
