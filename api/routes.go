package api

import (
	"github.com/labstack/echo/v5"

	"github.com/ajn2004/go-weather/weather"
)

func RegisterRoutes(e *echo.Echo, svc *weather.Service) {
	h := NewHandler(svc)

	e.GET("/health", h.Health)

	e.GET("/weather", h.GetWeather)
	e.POST("/weather", h.RefreshWeather)

	e.GET("/weather/current", h.GetCurrentWeather)
	e.POST("/weather/current", h.RefreshCurrent)

	e.GET("/weather/hourly", h.GetHourlyForecast)
	e.POST("/weather/hourly", h.RefreshHourly)

	e.GET("/weather/daily", h.GetDailyForecast)
	e.POST("/weather/daily", h.RefreshDaily)
}
