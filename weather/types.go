// weather/types.go
package weather

import (
	"time"
)

// We're defining our weather types that will be used for our app

// app side types
type WindDirection string

const (
	North     WindDirection = "N"
	Northeast WindDirection = "NE"
	East      WindDirection = "E"
	Southeast WindDirection = "SE"
	South     WindDirection = "S"
	Southwest WindDirection = "SW"
	West      WindDirection = "W"
	Northwest WindDirection = "NW"
	Variable  WindDirection = "VRBL"
	NA        WindDirection = "NA"
)

type CurrentWeather struct {
	Temperature   float64
	Humidity      float64
	WindSpeed     float64
	Pressure      float64
	DewPoint      float64
	Visibility    float64
	WindDirection WindDirection
	Condition     string
	LastUpdate    time.Time
}

type HourlyForecast struct {
	Time          time.Time
	Temperature   float64
	Humidity      float64
	WindSpeed     float64
	WindDirection WindDirection
	DewPoint      float64
	PrecipChance  float64
	Condition     string
}

type DailyForecast struct {
	Date         time.Time
	HighTemp     float64
	LowTemp      float64
	Humidity     float64
	Condition    string
	PrecipChance float64
}

type WeatherData struct {
	Current     CurrentWeather
	Hourly      []HourlyForecast
	Daily       []DailyForecast
	LastRefresh time.Time
}
