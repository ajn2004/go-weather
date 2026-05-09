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

}

func StartRefreshLoop(svc *weather.Service) {
	startRefreshLoopWithIntervals(svc, 5*time.Minute, 15*time.Minute, nil)
}

func startRefreshLoopWithIntervals(svc *weather.Service, currentInterval, fullInterval time.Duration, stop <-chan struct{}) {
	go func() {
		currentTicker := time.NewTicker(currentInterval)
		fullTicker := time.NewTicker(fullInterval)
		defer currentTicker.Stop()
		defer fullTicker.Stop()

		for {
			select {
			case <-stop:
				return
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
