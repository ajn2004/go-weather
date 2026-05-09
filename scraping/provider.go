package scraping

import "github.com/ajn2004/go-weather/weather"

// Weather Provider
type WeatherGovProvider struct{}

func (p WeatherGovProvider) GetCurrentObservation(stationID string) (weather.CurrentWeather, error) {
	return GetCurrentObservation(stationID)
}

func (p WeatherGovProvider) GetHourlyForecast(forecastOfficeID string, forecastGridX int, forecastGridY int) ([]weather.HourlyForecast, error) {
	return GetHourlyForecast(forecastOfficeID, forecastGridX, forecastGridY)
}

func (p WeatherGovProvider) GetDailyForecast(forecastOfficeID string, forecastGridX int, forecastGridY int) ([]weather.DailyForecast, error) {
	return GetDailyForecast(forecastOfficeID, forecastGridX, forecastGridY)
}
