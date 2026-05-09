package main

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/ajn2004/go-weather/weather"
)

type schedulingProvider struct {
	currentCalls atomic.Int64
	hourlyCalls  atomic.Int64
	dailyCalls   atomic.Int64
}

func (p *schedulingProvider) GetCurrentObservation(_ string) (weather.CurrentWeather, error) {
	p.currentCalls.Add(1)
	return weather.CurrentWeather{}, nil
}

func (p *schedulingProvider) GetHourlyForecast(_ string, _, _ int) ([]weather.HourlyForecast, error) {
	p.hourlyCalls.Add(1)
	return []weather.HourlyForecast{{}}, nil
}

func (p *schedulingProvider) GetDailyForecast(_ string, _, _ int) ([]weather.DailyForecast, error) {
	p.dailyCalls.Add(1)
	return []weather.DailyForecast{{}}, nil
}

func TestStartRefreshLoopWithIntervals_TriggersCurrentAndFullRefresh(t *testing.T) {
	p := &schedulingProvider{}
	svc := weather.NewService(p, "KNYC", "OKX", 34, 48)
	stop := make(chan struct{})
	startRefreshLoopWithIntervals(svc, 10*time.Millisecond, 25*time.Millisecond, stop)

	time.Sleep(70 * time.Millisecond)
	close(stop)
	time.Sleep(5 * time.Millisecond)

	if p.currentCalls.Load() == 0 {
		t.Fatalf("expected current refresh calls > 0")
	}
	if p.hourlyCalls.Load() == 0 || p.dailyCalls.Load() == 0 {
		t.Fatalf("expected full refresh calls > 0")
	}
}
