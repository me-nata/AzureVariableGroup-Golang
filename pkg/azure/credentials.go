package azure

import (
	"github.com/joho/godotenv"
	"os"
)

type AzureCredentials struct {
    Pat string
    Org string
}

func GetCredentials() (string, string) {
    godotenv.Load()
    
    PAT := os.Getenv("PERSONAL_ACCESS_TOKEN")
    ORG := os.Getenv("ORGANIZATION_NAME")

    return PAT, ORG
}