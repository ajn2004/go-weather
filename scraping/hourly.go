package scraping

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ajn2004/go-weather/scraping/observations"
	"github.com/ajn2004/go-weather/weather"
)

func GetHourlyForecast(forecastOfficeID string, forecastGridX int, forecastGridY int) ([]weather.HourlyForecast, error) {
	var forecast []weather.HourlyForecast

	url := fmt.Sprintf("https://api.weather.gov/gridpoints/%s/%d,%d/forecast/hourly", forecastOfficeID, forecastGridX, forecastGridY)
	client := getClient()

	req, err := getRequest(url)
	if err != nil {
		return forecast, err
	}

	res, err := client.Do(req)
	if err != nil {
		return forecast, fmt.Errorf("error making hourly request: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return forecast, fmt.Errorf("unexpected status code while fetching hourly: %d", res.StatusCode)
	}
	defer res.Body.Close()

	var resp observations.HourlyForecastResponse

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return forecast, fmt.Errorf("error decoding hourly forecast response: %w", err)
	}

	forecast = mapHourlyResponseToHourlyForecast(resp, 12)
	return forecast, nil
}
