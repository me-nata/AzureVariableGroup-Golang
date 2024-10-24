package azure

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AzureCredentials struct {
    Pat string
    Org string
}

func GetCredentials() (AzureCredentials, error) {
	// Load environment variables from .env file
    err := godotenv.Load()
    
	if err != nil {
		log.Fatal("Error loading .env file")
		return AzureCredentials{}, err
	}
	
	PAT := os.Getenv("PERSONAL_ACCESS_TOKEN")
	ORG := os.Getenv("ORGANIZATION_NAME")

    return AzureCredentials{ PAT, ORG }, nil
}