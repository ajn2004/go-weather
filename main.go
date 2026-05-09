package main

import (
	"fmt"

	"github.com/ajn2004/go-weather/scraping"
)

func main() {

	provider := scraping.WeatherGovProvider{}

	obs, _ := provider.GetCurrentObservation("KNYC")
	fmt.Printf("Observation: %v\n", obs)

	forecast, _ := provider.GetHourlyForecast("OKX", 34, 48)
	fmt.Printf("Forecast: %v\n", forecast[0])

	daily, _ := provider.GetDailyForecast("OKX", 34, 48)
	fmt.Printf("Daily Forecast: %v\n", daily)
}
