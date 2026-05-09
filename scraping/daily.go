package scraping

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ajn2004/go-weather/scraping/observations"
	"github.com/ajn2004/go-weather/weather"
)

func GetDailyForecast(forecastOfficeID string, forecastGridX int, forecastGridY int) ([]weather.DailyForecast, error) {
	var forecast []weather.DailyForecast

	url := fmt.Sprintf("https://api.weather.gov/gridpoints/%s/%d,%d/forecast", forecastOfficeID, forecastGridX, forecastGridY)
	client := getClient()

	req, err := getRequest(url)
	if err != nil {
		return forecast, err
	}

	res, err := client.Do(req)
	if err != nil {
		return forecast, fmt.Errorf("error making daily request: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return forecast, fmt.Errorf("Unexpected status code while fetching daily: %d", res.StatusCode)
	}
	defer res.Body.Close()

	var resp observations.DailyForecastResponse
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return forecast, fmt.Errorf("error decoding daily forecast response: %w", err)
	}

	forecast = mapDailyResponseToDailyForecast(resp, 7)
	return forecast, nil
}
