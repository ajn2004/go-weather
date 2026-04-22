package scraping

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ajn2004/go-weather/scraping/observations"
	"github.com/ajn2004/go-weather/weather"
)

func GetCurrentObservation(stationID string) (weather.CurrentWeather, error) {
	var obs weather.CurrentWeather

	// Define the base URL for the API endpoint, inserting the station ID
	baseURL := "https://api.weather.gov/stations/" + stationID + "/observations/latest"
	// Create an HTTP client with a timeout to avoid hanging requests
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return obs, err
	}
	// Set a custom User-Agent header to identify the application and provide contact information in compliance with API usage guidelines
	req.Header.Set("User-Agent", "go-weather-scraper/1.0 (contact: "+os.Getenv("EMAIL")+")")
	req.Header.Set("Accept", "application/geo+json")

	// Make the request and check for errors
	res, err := client.Do(req)
	if res.StatusCode != http.StatusOK {
		return obs, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	defer res.Body.Close() // ensure the response body is closed after we're done with it

	// Decode the JSON response into the ObservationResponse struct
	var resp observations.ObservationResponse

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return obs, fmt.Errorf("error decoding response: %w", err)
	}

	fmt.Println("Date:", res.Header.Get("Date"))
	fmt.Println("Age:", res.Header.Get("Age"))
	fmt.Println("Cache-Control:", res.Header.Get("Cache-Control"))
	fmt.Println("Last-Modified:", res.Header.Get("Last-Modified"))
	fmt.Println("Expires:", res.Header.Get("Expires"))
	obs = mapObservationToCurrentWeather(resp)

	return obs, nil
}

func GetHourlyForecast(forecastOfficeId string, forecastGridX int, forecastGridY int) ([]weather.HourlyForecast, error) {
	var forecast []weather.HourlyForecast

	url := fmt.Sprintf("https://api.weather.gov/gridpoints/%s/%d,%d/forecast/hourly", forecastOfficeId, forecastGridX, forecastGridY)
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return forecast, err
	}

	req.Header.Set("User-Agent", "go-weather-scraper/1.0 (contact:"+os.Getenv("EMAIL")+")")
	req.Header.Set("Accept", "application/geo+json")

	res, err := client.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return forecast, fmt.Errorf("error fetching hourly forecast: %w", err)
	}
	defer res.Body.Close()

	var resp observations.HourlyForecastResponse

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return forecast, fmt.Errorf("error decoding hourly forecast response: %w", err)
	}

	forecast = mapHourlyResponseToHourlyForecast(resp, 12)
	return forecast, nil
}
