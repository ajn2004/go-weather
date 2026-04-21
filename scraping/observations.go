// scraping/observations.go
package scraping

import "time"

// api.weather.gov response types

// ObservationResponse represents the JSON response from the /stations/{stationId}/observations/latest endpoint
type ObservationResponse struct {
	Properties struct {
		StationName     string    `json:"stationName"`
		Timestamp       time.Time `json:"timestamp"`
		Condition       string    `json:"textDescription"`
		RawMessage      string    `json:"rawMessage"`
		TextDescription string    `json:"textDescription"`
		iconURL         string    `json:"icon"`

		Temperature struct {
			Value          *float64 `json:"value"`
			Unit           string   `json:"unitCode"`
			QualityControl string   `json:"qualityControl"`
		} `json:"temperature"`
		DewPoint struct {
			Value          *float64 `json:"value"`
			Unit           string   `json:"unitCode"`
			QualityControl string   `json:"qualityControl"`
		} `json:"dewpoint"`
		WindDirection struct {
			Value          *float64 `json:"value"`
			Unit           string   `json:"unitCode"`
			QualityControl string   `json:"qualityControl"`
		} `json:"windDirection"`
		WindSpeed struct {
			Value          *float64 `json:"value"`
			Unit           string   `json:"unitCode"`
			QualityControl string   `json:"qualityControl"`
		} `json:"windSpeed"`
		WindGust struct {
			Value          *float64 `json:"value"`
			Unit           string   `json:"unitCode"`
			QualityControl string   `json:"qualityControl"`
		} `json:"windGust"`
		Pressure struct {
			Value          *float64 `json:"value"`
			Unit           string   `json:"unitCode"`
			QualityControl string   `json:"qualityControl"`
		} `json:"barometricPressure"`
		SeaLevelPressure struct {
			Value          *float64 `json:"value"`
			Unit           string   `json:"unitCode"`
			QualityControl string   `json:"qualityControl"`
		} `json:"seaLevelPressure"`
		Visibility struct {
			Value          *float64 `json:"value"`
			Unit           string   `json:"unitCode"`
			QualityControl string   `json:"qualityControl"`
		} `json:"visibility"`
		MaxTemperatureLast24Hours struct {
			Value          *float64 `json:"value"`
			Unit           string   `json:"unitCode"`
			QualityControl string   `json:"qualityControl"`
		} `json:"maxTemperatureLast24Hours"`
		MinTemperatureLast24Hours struct {
			Value          *float64 `json:"value"`
			Unit           string   `json:"unitCode"`
			QualityControl string   `json:"qualityControl"`
		} `json:"minTemperatureLast24Hours"`
		PrecipitationLastHour struct {
			Value          *float64 `json:"value"`
			Unit           string   `json:"unitCode"`
			QualityControl string   `json:"qualityControl"`
		} `json:"precipitationLastHour"`
		PrecipitationLast3Hours struct {
			Value          *float64 `json:"value"`
			Unit           string   `json:"unitCode"`
			QualityControl string   `json:"qualityControl"`
		} `json:"precipitationLast3Hours"`
		PrecipitationLast6Hours struct {
			Value          *float64 `json:"value"`
			Unit           string   `json:"unitCode"`
			QualityControl string   `json:"qualityControl"`
		} `json:"precipitationLast6Hours"`
		Humidity struct {
			Value          *float64 `json:"value"`
			Unit           string   `json:"unitCode"`
			QualityControl string   `json:"qualityControl"`
		} `json:"relativeHumidity"`
		WindChill struct {
			Value          *float64 `json:"value"`
			Unit           string   `json:"unitCode"`
			QualityControl string   `json:"qualityControl"`
		} `json:"windChill"`
		HeatIndex struct {
			Value          *float64 `json:"value"`
			Unit           string   `json:"unitCode"`
			QualityControl string   `json:"qualityControl"`
		} `json:"heatIndex"`
	} `json:"properties"`
}

// HourForecastResponse represents the JSON response from the /gridpoints/{office}/{gridX},{gridY}/forecast/hourly endpoint
type HourForecastResponse struct {
	Properties struct {
		periods []struct {
			Number              int       `json:"number"`
			Name                string    `json:"name"`
			StartTime           time.Time `json:"startTime"`
			EndTime             time.Time `json:"endTime"`
			isDaytime           bool      `json:"isDaytime"`
			Temperature         float64   `json:"temperature"`
			TemperatureUnit     string    `json:"temperatureUnit"`
			TempreatureTrend    *string   `json:"temperatureTrend,omitempty"`
			PrecipitationChance struct {
				Value *float64 `json:"value"`
				Unit  string   `json:"unitCode"`
			} `json:"probabilityOfPrecipitation"`
			DewPoint struct {
				Value *float64 `json:"value"`
				Unit  string   `json:"unitCode"`
			} `json:"dewpoint"`
			Humidity struct {
				Value *float64 `json:"value"`
				Unit  string   `json:"unitCode"`
			} `json:"relativeHumidity"`
			WindSpeed        string `json:"windSpeed"`
			WindDirection    string `json:"windDirection"`
			IconURL          string `json:"icon"`
			ShortForecast    string `json:"shortForecast"`
			DetailedForecast string `json:"detailedForecast"`
		} `json:"periods"`
	} `json:"properties"`
	GenerationTime time.Time `json:"generationTime"`
	UpdateTime     time.Time `json:"updateTime"`
	ValidTimes     time.Time `json:"validTimes"`
}

// DailyForecastResponse represents the JSON response from the /gridpoints/{office}/{gridX},{gridY}/forecast endpoint
type DailyForecastResponse struct {
	Properties struct {
		periods []struct {
			Number              int       `json:"number"`
			Name                string    `json:"name"`
			StartTime           time.Time `json:"startTime"`
			EndTime             time.Time `json:"endTime"`
			isDaytime           bool      `json:"isDaytime"`
			Temperature         float64   `json:"temperature"`
			TemperatureUnit     string    `json:"temperatureUnit"`
			TempreatureTrend    *string   `json:"temperatureTrend,omitempty"`
			PrecipitationChance struct {
				Value *float64 `json:"value"`
				Unit  string   `json:"unitCode"`
			} `json:"probabilityOfPrecipitation"`
			WindSpeed        string `json:"windSpeed"`
			WindDirection    string `json:"windDirection"`
			IconURL          string `json:"icon"`
			ShortForecast    string `json:"shortForecast"`
			DetailedForecast string `json:"detailedForecast"`
		} `json:"periods"`
	} `json:"properties"`
	GenerationTime time.Time `json:"generationTime"`
	UpdateTime     time.Time `json:"updateTime"`
	ValidTimes     time.Time `json:"validTimes"`
}
