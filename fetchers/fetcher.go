package fetchers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// FetchJSON fetches raw JSON data from a URL and returns it as bytes.
func FetchJSON(url string) ([]byte, error) {
	client := http.Client{
		Timeout: time.Second * 10, // prevent hanging requests
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}

	// Use GitHub token if available
	token := os.Getenv("GITHUB_TOKEN")
	if token != "" {
		req.Header.Set("Authorization", "token "+token)
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return body, nil
}
