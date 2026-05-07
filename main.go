package main

import (
	"fmt"

	"github.com/ajn2004/go-weather/scraping"
)

func main() {
	obs, _ := scraping.GetCurrentObservation("KNYC")
	fmt.Printf("Observation: %v\n", obs)

	forecast, _ := scraping.GetHourlyForecast("OKX", 34, 48)
	fmt.Printf("Forecast: %v\n", forecast[0])

	daily, _ := scraping.GetDailyForecast("OKX", 34, 48)
	fmt.Printf("Daily Forecast: %v\n", daily)
}
