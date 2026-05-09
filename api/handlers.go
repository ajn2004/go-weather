package api

import (
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/ajn2004/go-weather/weather"
)

type Handler struct {
	service *weather.Service
}

func NewHandler(service *weather.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Health(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) GetWeather(c *echo.Context) error {
	return c.JSON(http.StatusOK, h.service.GetWeatherData())
}

func (h *Handler) GetCurrentWeather(c *echo.Context) error {
	return c.JSON(http.StatusOK, h.service.GetCurrent())
}

func (h *Handler) GetHourlyForecast(c *echo.Context) error {
	return c.JSON(http.StatusOK, h.service.GetHourly())
}

func (h *Handler) GetDailyForecast(c *echo.Context) error {
	return c.JSON(http.StatusOK, h.service.GetDaily())
}

func (h *Handler) RefreshWeather(c *echo.Context) error {
	if err := h.service.Refresh(); err != nil {
		return errorJSON(c, err, http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, h.service.GetWeatherData())
}

func (h *Handler) RefreshCurrent(c *echo.Context) error {
	if err := h.service.RefreshCurrent(); err != nil {
		return errorJSON(c, err, http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, h.service.GetCurrent())
}

func (h *Handler) RefreshDaily(c *echo.Context) error {
	if err := h.service.RefreshDaily(); err != nil {
		return errorJSON(c, err, http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, h.service.GetDaily())
}

func (h *Handler) RefreshHourly(c *echo.Context) error {
	if err := h.service.RefreshHourly(); err != nil {
		return errorJSON(c, err, http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, h.service.GetHourly())
}

func errorJSON(c *echo.Context, err error, status int) error {
	return c.JSON(status, map[string]string{"error": err.Error()})
}
