package azureapi

import (
	"fmt"
	"io"
	"net/http"

	"github.com/me-nata/goazuredevops/pkg/azure"
)

// AzureApi represents the Azure API structure holding credentials
type AzureApiCaller struct {
    Credentials *azure.AzureCredentials
}

// Makes a GET request to the given URL and returns the response body as []byte
func (api *AzureApiCaller) get(url string) ([]byte, error) {
    client := &http.Client{}

    // Create the GET request
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to create GET request: %w", err)
    }

    // Set request headers
    req.Header.Add("Accept", "application/json")
    req.Header.Add("Authorization", basicAuth(api.Credentials.Pat))

    // Execute the request
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("failed to execute GET request: %w", err)
    }
    defer resp.Body.Close()

    // Check if status code is not 200 OK
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected status code: %d, expected: 200", resp.StatusCode)
    }

    // Read response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read response body: %w", err)
    }

    return body, nil
}

func basicAuth(token string) string {
    return "Basic " + token // Example, change as needed
}