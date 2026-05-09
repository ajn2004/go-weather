package scraping

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) { return f(req) }

func jsonResp(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(bytes.NewBufferString(body)),
		Header:     make(http.Header),
	}
}

func TestGetRequest_SetsHeaders(t *testing.T) {
	t.Setenv("EMAIL", "test@example.com")
	req, err := getRequest("https://api.weather.gov/test")
	if err != nil {
		t.Fatalf("getRequest failed: %v", err)
	}
	if got := req.Header.Get("Accept"); got != "application/geo+json" {
		t.Fatalf("unexpected Accept header: %q", got)
	}
	if got := req.Header.Get("User-Agent"); !strings.Contains(got, "test@example.com") {
		t.Fatalf("expected User-Agent to include EMAIL, got %q", got)
	}
}

func TestGetCurrentObservation_Success(t *testing.T) {
	t.Setenv("EMAIL", "test@example.com")
	oldTransport := http.DefaultTransport
	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		if !strings.Contains(req.URL.String(), "/stations/KNYC/observations/latest") {
			t.Fatalf("unexpected URL: %s", req.URL.String())
		}
		body := `{"properties":{"timestamp":"2026-01-01T00:00:00Z","textDescription":"Clear","temperature":{"value":20.0,"unitCode":"wmoUnit:degC"},"relativeHumidity":{"value":50.0},"windSpeed":{"value":5.0,"unitCode":"wmoUnit:km_h-1"},"barometricPressure":{"value":101325.0,"unitCode":"wmoUnit:Pa"},"dewpoint":{"value":5.0,"unitCode":"wmoUnit:degC"},"visibility":{"value":1000.0,"unitCode":"wmoUnit:m"},"windDirection":{"value":90.0}}}`
		return jsonResp(http.StatusOK, body), nil
	})
	t.Cleanup(func() { http.DefaultTransport = oldTransport })

	got, err := GetCurrentObservation("KNYC")
	if err != nil {
		t.Fatalf("GetCurrentObservation failed: %v", err)
	}
	if got.Condition != "Clear" {
		t.Fatalf("expected Clear condition, got %q", got.Condition)
	}
	if got.WindDirection == "" {
		t.Fatalf("expected mapped wind direction")
	}
}

func TestGetHourlyForecast_HTTPError(t *testing.T) {
	oldTransport := http.DefaultTransport
	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		return jsonResp(http.StatusBadGateway, `{"error":"bad gateway"}`), nil
	})
	t.Cleanup(func() { http.DefaultTransport = oldTransport })

	_, err := GetHourlyForecast("OKX", 34, 48)
	if err == nil {
		t.Fatalf("expected hourly error for non-200 response")
	}
}

func TestGetDailyForecast_Success(t *testing.T) {
	oldTransport := http.DefaultTransport
	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		if !strings.Contains(req.URL.String(), "/gridpoints/OKX/34,48/forecast") {
			t.Fatalf("unexpected URL: %s", req.URL.String())
		}
		body := `{"properties":{"periods":[{"name":"Today","startTime":"2026-01-01T00:00:00Z","endTime":"2026-01-01T12:00:00Z","isDaytime":true,"temperature":20,"temperatureUnit":"C","probabilityOfPrecipitation":{"value":10},"windSpeed":"10 mph","windDirection":"N","shortForecast":"Sunny"}]}}`
		return jsonResp(http.StatusOK, body), nil
	})
	t.Cleanup(func() { http.DefaultTransport = oldTransport })

	got, err := GetDailyForecast("OKX", 34, 48)
	if err != nil {
		t.Fatalf("GetDailyForecast failed: %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("expected 1 daily period, got %d", len(got))
	}
}
