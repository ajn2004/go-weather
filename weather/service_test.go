package weather

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type mockProvider struct {
	currentErr error
	hourlyErr  error
	dailyErr   error

	currentCalls atomic.Int64
	hourlyCalls  atomic.Int64
	dailyCalls   atomic.Int64
}

func (m *mockProvider) GetCurrentObservation(_ string) (CurrentWeather, error) {
	m.currentCalls.Add(1)
	if m.currentErr != nil {
		return CurrentWeather{}, m.currentErr
	}
	return CurrentWeather{Temperature: 20, Condition: "Clear", LastUpdate: time.Now()}, nil
}

func (m *mockProvider) GetHourlyForecast(_ string, _, _ int) ([]HourlyForecast, error) {
	m.hourlyCalls.Add(1)
	if m.hourlyErr != nil {
		return nil, m.hourlyErr
	}
	return []HourlyForecast{{Temperature: 21, Condition: "Sunny"}}, nil
}

func (m *mockProvider) GetDailyForecast(_ string, _, _ int) ([]DailyForecast, error) {
	m.dailyCalls.Add(1)
	if m.dailyErr != nil {
		return nil, m.dailyErr
	}
	return []DailyForecast{{Name: "Today", Temperature: 22}}, nil
}

func TestServiceRefreshAndGetters(t *testing.T) {
	p := &mockProvider{}
	svc := NewService(p, "KNYC", "OKX", 34, 48)

	if err := svc.Refresh(); err != nil {
		t.Fatalf("refresh failed: %v", err)
	}

	if svc.GetCurrent().Condition != "Clear" {
		t.Fatalf("expected current condition Clear")
	}
	if len(svc.GetHourly()) != 1 {
		t.Fatalf("expected hourly data")
	}
	if len(svc.GetDaily()) != 1 {
		t.Fatalf("expected daily data")
	}
}

func TestServiceRefreshErrorPropagation(t *testing.T) {
	p := &mockProvider{hourlyErr: errors.New("hourly fail")}
	svc := NewService(p, "KNYC", "OKX", 34, 48)

	if err := svc.Refresh(); err == nil {
		t.Fatalf("expected refresh error")
	}
}

func TestServiceConcurrentAccess_NoRacePattern(t *testing.T) {
	p := &mockProvider{}
	svc := NewService(p, "KNYC", "OKX", 34, 48)
	if err := svc.Refresh(); err != nil {
		t.Fatalf("seed refresh failed: %v", err)
	}

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 200; j++ {
				_ = svc.RefreshCurrent()
				_ = svc.RefreshHourly()
				_ = svc.RefreshDaily()
			}
		}()
	}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 500; j++ {
				_ = svc.GetCurrent()
				_ = svc.GetHourly()
				_ = svc.GetDaily()
				_ = svc.GetWeatherData()
			}
		}()
	}

	wg.Wait()
	if p.currentCalls.Load() == 0 || p.hourlyCalls.Load() == 0 || p.dailyCalls.Load() == 0 {
		t.Fatalf("expected provider calls to occur")
	}
}

func BenchmarkServiceGetWeatherData(b *testing.B) {
	p := &mockProvider{}
	svc := NewService(p, "KNYC", "OKX", 34, 48)
	_ = svc.Refresh()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = svc.GetWeatherData()
	}
}
