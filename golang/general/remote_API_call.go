package general

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// GitHubAPIInfo holds some fields from the GitHub API root response
type GitHubAPIInfo struct {
	CurrentUserURL    string `json:"current_user_url"`
	AuthorizationsURL string `json:"authorizations_url"`
}

func remoteAPICall() {
	url := "https://api.github.com"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	fmt.Println("Response:")
	fmt.Println(string(body))

	decoder := json.NewDecoder(resp.Body)

	var apiInfo GitHubAPIInfo
	if err := decoder.Decode(&apiInfo); err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}

	fmt.Println("Decoded Response:")
	fmt.Printf("Current User URL: %s\n", apiInfo.CurrentUserURL)
	fmt.Printf("Authorizations URL: %s\n", apiInfo.AuthorizationsURL)
}
