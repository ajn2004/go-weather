package observations

import "time"

// ObservationResponse represents the JSON response from the /stations/{stationId}/observations/latest endpoint
type ObservationResponse struct {
	ID         string                `json:"id,omitempty"`
	Type       string                `json:"type,omitempty"` // GeoJSON feature
	Geometry   any                   `json:"geometry,omitempty"`
	Properties ObservationProperties `json:"properties"`
}

type ObservationProperties struct {
	ID              string    `json:"@id,omitempty"`
	Station         string    `json:"station,omitempty"`
	StationID       string    `json:"stationId,omitempty"`
	StationName     string    `json:"stationName,omitempty"`
	Timestamp       time.Time `json:"timestamp"`
	RawMessage      string    `json:"rawMessage,omitempty"`
	TextDescription string    `json:"textDescription,omitempty"`
	Icon            *string   `json:"icon,omitempty"`

	Temperature               *QuantitativeValue `json:"temperature,omitempty"`
	Dewpoint                  *QuantitativeValue `json:"dewpoint,omitempty"`
	WindDirection             *QuantitativeValue `json:"windDirection,omitempty"`
	WindSpeed                 *QuantitativeValue `json:"windSpeed,omitempty"`
	WindGust                  *QuantitativeValue `json:"windGust,omitempty"`
	BarometricPressure        *QuantitativeValue `json:"barometricPressure,omitempty"`
	SeaLevelPressure          *QuantitativeValue `json:"seaLevelPressure,omitempty"`
	Visibility                *QuantitativeValue `json:"visibility,omitempty"`
	MaxTemperatureLast24Hours *QuantitativeValue `json:"maxTemperatureLast24Hours,omitempty"`
	MinTemperatureLast24Hours *QuantitativeValue `json:"minTemperatureLast24Hours,omitempty"`
	PrecipitationLastHour     *QuantitativeValue `json:"precipitationLastHour,omitempty"`
	PrecipitationLast3Hours   *QuantitativeValue `json:"precipitationLast3Hours,omitempty"`
	PrecipitationLast6Hours   *QuantitativeValue `json:"precipitationLast6Hours,omitempty"`
	RelativeHumidity          *QuantitativeValue `json:"relativeHumidity,omitempty"`
	WindChill                 *QuantitativeValue `json:"windChill,omitempty"`
	HeatIndex                 *QuantitativeValue `json:"heatIndex,omitempty"`

	CloudLayers []CloudLayer `json:"cloudLayers,omitempty"`
}

type CloudLayer struct {
	Base   *QuantitativeValue `json:"base,omitempty"`
	Amount string             `json:"amount,omitempty"`
}
