package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/ajn2004/go-weather/api"
	"github.com/ajn2004/go-weather/scraping"
	"github.com/ajn2004/go-weather/weather"
	"github.com/labstack/echo/v5"
)

func main() {

	// Example usage of the WeatherGovProvider
	provider := scraping.WeatherGovProvider{}
	gridx, err := strconv.Atoi(os.Getenv("GRIDX"))
	if err != nil {
		log.Fatalf("Invalid GRIDX value: %v", err)
	}
	gridy, err := strconv.Atoi(os.Getenv("GRIDY"))
	if err != nil {
		log.Fatalf("Invalid GRIDY value: %v", err)
	}
	svc := weather.NewService(provider, os.Getenv("STATION"), os.Getenv("OFFICE"), gridx, gridy)

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
