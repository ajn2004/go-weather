package scraping

import (
	"testing"
	"time"

	"github.com/ajn2004/go-weather/scraping/observations"
	"github.com/ajn2004/go-weather/weather"
)

func f64ptr(v float64) *float64 { return &v }

func TestMapObservationToCurrentWeather_ConvertsAndMaps(t *testing.T) {
	now := time.Now()
	resp := observations.ObservationResponse{
		Properties: observations.ObservationProperties{
			Timestamp:       now,
			TextDescription: "Clear",
			Temperature:     &observations.QuantitativeValue{Value: f64ptr(86), UnitCode: "wmoUnit:degF"},
			RelativeHumidity: &observations.QuantitativeValue{
				Value: f64ptr(65),
			},
			WindSpeed:          &observations.QuantitativeValue{Value: f64ptr(18), UnitCode: "wmoUnit:km_h-1"},
			BarometricPressure: &observations.QuantitativeValue{Value: f64ptr(1013.25), UnitCode: "wmoUnit:hPa"},
			Dewpoint:           &observations.QuantitativeValue{Value: f64ptr(10), UnitCode: "wmoUnit:degC"},
			Visibility:         &observations.QuantitativeValue{Value: f64ptr(1), UnitCode: "wmoUnit:mi"},
			WindDirection:      &observations.QuantitativeValue{Value: f64ptr(90)},
		},
	}

	got := mapObservationToCurrentWeather(resp)

	if got.Temperature < 29.9 || got.Temperature > 30.1 {
		t.Fatalf("expected ~30C, got %v", got.Temperature)
	}
	if got.WindSpeed < 4.9 || got.WindSpeed > 5.1 {
		t.Fatalf("expected ~5 m/s, got %v", got.WindSpeed)
	}
	if got.Pressure < 0.99 || got.Pressure > 1.01 {
		t.Fatalf("expected ~1 atm, got %v", got.Pressure)
	}
	if got.Visibility < 1609 || got.Visibility > 1610 {
		t.Fatalf("expected ~1609.34 meters visibility, got %v", got.Visibility)
	}
	if got.WindDirection != weather.East {
		t.Fatalf("expected east wind, got %s", got.WindDirection)
	}
	if got.Condition != "Clear" {
		t.Fatalf("expected condition Clear, got %q", got.Condition)
	}
	if !got.LastUpdate.Equal(now) {
		t.Fatalf("expected timestamp preserved")
	}
}

func TestMapHourlyResponseToHourlyForecast_MixedShapes(t *testing.T) {
	icon := "https://example/icon.png"
	tempInt := 72
	resp := observations.HourlyForecastResponse{
		Properties: observations.HourlyForecastProperties{
			Periods: []observations.HourlyForecastPeriod{
				{
					StartTime:       time.Now(),
					ShortForecast:   "Sunny",
					Icon:            &icon,
					Temperature:     observations.TemperatureValue{Int: &tempInt},
					TemperatureUnit: "F",
					WindSpeed:       observations.WindSpeedValue{Str: strptr("10 mph")},
					WindDirection:   "SW",
					RelativeHumidity: &observations.QuantitativeValue{
						Value: f64ptr(40),
					},
					Dewpoint:                   &observations.QuantitativeValue{Value: f64ptr(5), UnitCode: "C"},
					ProbabilityOfPrecipitation: &observations.QuantitativeValue{Value: f64ptr(30)},
				},
			},
		},
	}

	got := mapHourlyResponseToHourlyForecast(resp, 1)
	if len(got) != 1 {
		t.Fatalf("expected 1 period, got %d", len(got))
	}
	if got[0].Temperature < 22 || got[0].Temperature > 23 {
		t.Fatalf("expected temperature converted to C, got %v", got[0].Temperature)
	}
	if got[0].WindSpeed != 10 {
		t.Fatalf("expected raw wind speed 10 for string mph input, got %v", got[0].WindSpeed)
	}
	if got[0].WindDirection != weather.Southwest {
		t.Fatalf("expected southwest wind, got %s", got[0].WindDirection)
	}
	if got[0].PrecipChance != 30 {
		t.Fatalf("expected precip 30, got %d", got[0].PrecipChance)
	}
}

func TestMapDailyResponseToDailyForecast_LimitsToRequestedCount(t *testing.T) {
	temp := 20
	periods := []observations.DailyForecastPeriod{
		{Temperature: observations.TemperatureValue{Int: &temp}, TemperatureUnit: "C", ProbabilityOfPrecipitation: &observations.QuantitativeValue{Value: f64ptr(10)}},
		{Temperature: observations.TemperatureValue{Int: &temp}, TemperatureUnit: "C", ProbabilityOfPrecipitation: &observations.QuantitativeValue{Value: f64ptr(20)}},
	}

	got := mapDailyResponseToDailyForecast(observations.DailyForecastResponse{
		Properties: observations.DailyForecastProperties{Periods: periods},
	}, 1)
	if len(got) != 1 {
		t.Fatalf("expected 1 period, got %d", len(got))
	}
}

func BenchmarkMapHourlyResponseToHourlyForecast(b *testing.B) {
	temp := 70
	resp := observations.HourlyForecastResponse{
		Properties: observations.HourlyForecastProperties{
			Periods: []observations.HourlyForecastPeriod{
				{
					Temperature:     observations.TemperatureValue{Int: &temp},
					TemperatureUnit: "F",
					WindSpeed:       observations.WindSpeedValue{Str: strptr("10 mph")},
					WindDirection:   "N",
					ProbabilityOfPrecipitation: &observations.QuantitativeValue{
						Value: f64ptr(10),
					},
				},
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapHourlyResponseToHourlyForecast(resp, 1)
	}
}

func strptr(v string) *string { return &v }
