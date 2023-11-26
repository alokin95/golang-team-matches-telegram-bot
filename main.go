package main

import (
	"kadigramo/assistant"
	"kadigramo/client"
	"kadigramo/models"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := loadEnv(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	assistantApiKey := os.Getenv("ASSISTANT_API_KEY")
	httpClient := client.NewHttpClient(assistantApiKey)

	messages := []models.Message{
		{
			Role:    "user",
			Content: "Kad igramo?",
		},
	}

	thread, err := assistant.CreateThread(*httpClient, messages)
	if err != nil {
		log.Fatalf("Failed to create thread: %v", err)
	}

	log.Println(thread.ID)
}

func loadEnv() error {
	return godotenv.Load()
}
