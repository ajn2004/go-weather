package scraping

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func getClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

func getRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil,
			fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("User-Agent", "go-weather-scraper/1.0 (contact:"+os.Getenv("EMAIL")+")")
	req.Header.Set("Accept", "application/geo+json")
	return req, nil
}
