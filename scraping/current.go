package scraping

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ajn2004/go-weather/scraping/observations"
	"github.com/ajn2004/go-weather/weather"
)

// Function definitions
func GetCurrentObservation(stationID string) (weather.CurrentWeather, error) {
	var obs weather.CurrentWeather

	// Define the base URL for the API endpoint, inserting the station ID
	baseURL := "https://api.weather.gov/stations/" + stationID + "/observations/latest"
	// Create an HTTP client with a timeout to avoid hanging requests
	client := getClient()
	req, err := getRequest(baseURL)
	if err != nil {
		return obs, err
	}

	// Make the request and check for errors
	res, err := client.Do(req)
	if err != nil {
		return obs, fmt.Errorf("error making request: %w", err)
	}
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
