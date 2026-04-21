// scraping/observations.go
package scraping

import "time"

// api.weather.gov response types

type ObservationResponse struct {
	Properties struct {
		StationName string    `json:"stationName"`
		Timestamp   time.Time `json:"timestamp"`
		Condition   string    `json:"textDescription"`
		Temperature struct {
			Value *float64 `json:"value"`
			Unit  string   `json:"unitCode"`
		} `json:"temperature"`
		Humidity struct {
			Value *float64 `json:"value"`
			Unit  string   `json:"unitCode"`
		} `json:"relativeHumidity"`
		Pressure struct {
			Value *float64 `json:"value"`
			Unit  string   `json:"unitCode"`
		} `json:"barometricPressure"`
		DewPoint struct {
			Value *float64 `json:"value"`
			Unit  string   `json:"unitCode"`
		} `json:"dewpoint"`
		WindSpeed struct {
			Value *float64 `json:"value"`
			Unit  string   `json:"unitCode"`
		} `json:"windSpeed"`
		WindDirection struct {
			Value *float64 `json:"value"`
			Unit  string   `json:"unitCode"`
		} `json:"windDirection"`
		Visibility struct {
			Value *float64 `json:"value"`
			Unit  string   `json:"unitCode"`
		} `json:"visibility"`
	} `json:"properties"`
}
