package observations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// DailyForecastResponse represents the JSON response from the /gridpoints/{officeId}/{X},{Y}/forecast endpoint
type DailyForecastResponse struct {
	ID         string                  `json:"id,omitempty"`
	Type       string                  `json:"type,omitempty"` // GeoJSON feature
	Geometry   json.RawMessage         `json:"geometry,omitempty"`
	Properties DailyForecastProperties `json:"properties"`
}

type DailyForecastProperties struct {
	Units             string                `json:"units,omitempty"`
	ForecastGenerator string                `json:"forecastGenerator,omitempty"`
	GeneratedAt       time.Time             `json:"generatedAt,omitempty"`
	UpdateTime        time.Time             `json:"updateTime,omitempty"`
	ValidTimes        string                `json:"validTimes,omitempty"`
	Periods           []DailyForecastPeriod `json:"periods,omitempty"`
	Elevation         *QuantitativeValue    `json:"elevation,omitempty"`
}

type DailyForecastPeriod struct {
	Number                     int                `json:"number,omitempty"`
	Name                       string             `json:"name,omitempty"`
	StartTime                  time.Time          `json:"startTime,omitempty"`
	EndTime                    time.Time          `json:"endTime,omitempty"`
	IsDaytime                  bool               `json:"isDaytime,omitempty"`
	Temperature                TemperatureValue   `json:"temperature,omitempty"`
	TemperatureUnit            string             `json:"temperatureUnit,omitempty"` // deprecated, keep for compatibility
	TemperatureTrend           *string            `json:"temperatureTrend,omitempty"`
	ProbabilityOfPrecipitation *QuantitativeValue `json:"probabilityOfPrecipitation,omitempty"`
	WindSpeed                  WindSpeedValue     `json:"windSpeed,omitempty"`
	WindGust                   WindSpeedValue     `json:"windGust,omitempty"`
	WindDirection              string             `json:"windDirection,omitempty"`
	Icon                       *string            `json:"icon,omitempty"`
	ShortForecast              string             `json:"shortForecast,omitempty"`
	DetailedForecast           string             `json:"detailedForecast,omitempty"`
}

// HourlyForecastResponse represents the JSON response from the /gridpoints/{officeId}/{X},{Y}/forecast/hourly endpoint
type HourlyForecastResponse struct {
	ID         string                   `json:"id,omitempty"`
	Type       string                   `json:"type,omitempty"` // GeoJSON feature
	Geometry   json.RawMessage          `json:"geometry,omitempty"`
	Properties HourlyForecastProperties `json:"properties"`
}

type HourlyForecastProperties struct {
	Units             string                 `json:"units,omitempty"`
	ForecastGenerator string                 `json:"forecastGenerator,omitempty"`
	GeneratedAt       time.Time              `json:"generatedAt,omitempty"`
	UpdateTime        time.Time              `json:"updateTime,omitempty"`
	ValidTimes        string                 `json:"validTimes,omitempty"`
	Periods           []HourlyForecastPeriod `json:"periods,omitempty"`
	Elevation         *QuantitativeValue     `json:"elevation,omitempty"`
}

type HourlyForecastPeriod struct {
	Number                     int                `json:"number,omitempty"`
	Name                       string             `json:"name,omitempty"`
	StartTime                  time.Time          `json:"startTime,omitempty"`
	EndTime                    time.Time          `json:"endTime,omitempty"`
	IsDaytime                  bool               `json:"isDaytime,omitempty"`
	Temperature                TemperatureValue   `json:"temperature,omitempty"`
	TemperatureUnit            string             `json:"temperatureUnit,omitempty"` // deprecated, keep for compatibility
	TemperatureTrend           *string            `json:"temperatureTrend,omitempty"`
	ProbabilityOfPrecipitation *QuantitativeValue `json:"probabilityOfPrecipitation,omitempty"`
	Dewpoint                   *QuantitativeValue `json:"dewpoint,omitempty"`
	RelativeHumidity           *QuantitativeValue `json:"relativeHumidity,omitempty"`
	WindSpeed                  WindSpeedValue     `json:"windSpeed,omitempty"`
	WindGust                   WindSpeedValue     `json:"windGust,omitempty"`
	WindDirection              string             `json:"windDirection,omitempty"`
	Icon                       *string            `json:"icon,omitempty"`
	ShortForecast              string             `json:"shortForecast,omitempty"`
	DetailedForecast           string             `json:"detailedForecast,omitempty"`
}

func (t *TemperatureValue) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if bytes.Equal(data, []byte("null")) {
		*t = TemperatureValue{}
		return nil
	}
	var i int
	if err := json.Unmarshal(data, &i); err == nil {
		t.Int = &i
		t.QV = nil
		return nil
	}

	var qv QuantitativeValue
	if err := json.Unmarshal(data, &qv); err == nil {
		t.Int = nil
		t.QV = &qv
		return nil
	}

	return fmt.Errorf("temperature: unsupported JSON: %s", string(data))
}

func (w *WindSpeedValue) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if bytes.Equal(data, []byte("null")) {
		*w = WindSpeedValue{}
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		w.Str = &s
		w.QV = nil
		return nil
	}

	var qv QuantitativeValue
	if err := json.Unmarshal(data, &qv); err == nil {
		w.Str = nil
		w.QV = &qv
		return nil
	}

	return fmt.Errorf("wind speed: unsupported JSON: %s", string(data))
}
