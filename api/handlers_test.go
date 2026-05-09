package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v5"

	"github.com/ajn2004/go-weather/weather"
)

type handlerProvider struct {
	refreshErr error
}

func (p handlerProvider) GetCurrentObservation(_ string) (weather.CurrentWeather, error) {
	if p.refreshErr != nil {
		return weather.CurrentWeather{}, p.refreshErr
	}
	return weather.CurrentWeather{Condition: "Clear"}, nil
}

func (p handlerProvider) GetHourlyForecast(_ string, _, _ int) ([]weather.HourlyForecast, error) {
	if p.refreshErr != nil {
		return nil, p.refreshErr
	}
	return []weather.HourlyForecast{{Condition: "Sunny"}}, nil
}

func (p handlerProvider) GetDailyForecast(_ string, _, _ int) ([]weather.DailyForecast, error) {
	if p.refreshErr != nil {
		return nil, p.refreshErr
	}
	return []weather.DailyForecast{{Name: "Today"}}, nil
}

func setupEcho(t *testing.T, provider handlerProvider) *echo.Echo {
	t.Helper()
	svc := weather.NewService(provider, "KNYC", "OKX", 34, 48)
	if provider.refreshErr == nil {
		if err := svc.Refresh(); err != nil {
			t.Fatalf("seed refresh failed: %v", err)
		}
	}
	e := echo.New()
	RegisterRoutes(e, svc)
	return e
}

func TestHealthRoute(t *testing.T) {
	e := setupEcho(t, handlerProvider{})
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
}

func TestGetCurrentWeatherRoute(t *testing.T) {
	e := setupEcho(t, handlerProvider{})
	req := httptest.NewRequest(http.MethodGet, "/weather/current", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	var payload weather.CurrentWeather
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if payload.Condition != "Clear" {
		t.Fatalf("expected Clear condition, got %q", payload.Condition)
	}
}

func TestRefreshWeatherRoute_Error(t *testing.T) {
	e := setupEcho(t, handlerProvider{refreshErr: errors.New("provider error")})
	req := httptest.NewRequest(http.MethodPost, "/weather", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", rec.Code)
	}
}

func BenchmarkGetWeatherRoute(b *testing.B) {
	svc := weather.NewService(handlerProvider{}, "KNYC", "OKX", 34, 48)
	_ = svc.Refresh()
	e := echo.New()
	RegisterRoutes(e, svc)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/weather", nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			b.Fatalf("expected 200, got %d", rec.Code)
		}
	}
}
