package main

import (
	"log"
	"time"

	"github.com/ajn2004/go-weather/api"
	"github.com/ajn2004/go-weather/scraping"
	"github.com/ajn2004/go-weather/weather"
	"github.com/labstack/echo/v5"
)

func main() {

	// Example usage of the WeatherGovProvider
	provider := scraping.WeatherGovProvider{}

	svc := weather.NewService(provider, "KNYC", "OKX", 34, 48)

	if err := svc.Refresh(); err != nil {
		log.Fatalf("Failed to refresh weather data: %v", err)
	}

	StartRefreshLoop(svc)

	e := echo.New()

	api.RegisterRoutes(e, svc)

	log.Fatal(e.Start(":8080"))

	// obs, _ := provider.GetCurrentObservation("KNYC")
	// fmt.Printf("Observation: %v\n", obs)

	// forecast, _ := provider.GetHourlyForecast("OKX", 34, 48)
	// fmt.Printf("Forecast: %v\n", forecast[0])

	// daily, _ := provider.GetDailyForecast("OKX", 34, 48)
	// fmt.Printf("Daily Forecast: %v\n", daily)\

}

func StartRefreshLoop(svc *weather.Service) {
	go func() {
		currentTicker := time.NewTicker(5 * time.Minute)
		fullTicker := time.NewTicker(15 * time.Minute)
		defer currentTicker.Stop()
		defer fullTicker.Stop()

		for {
			select {
			case <-currentTicker.C:
				if err := svc.RefreshCurrent(); err != nil {
					log.Printf("Error refreshing current weather: %v", err)
				}
			case <-fullTicker.C:
				if err := svc.Refresh(); err != nil {
					log.Printf("Error refreshing full weather data: %v", err)
				}
			}
		}
	}()
}
