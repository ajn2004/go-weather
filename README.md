# go-weather

`go-weather` is a Go weather API service backed by NOAA/NWS (`api.weather.gov`).

It fetches current observations, hourly forecast, and daily forecast, normalizes units into app-facing types, caches data in memory, and serves it over HTTP.

## Current Features

- HTTP API built with Echo (`:8080`)
- Weather data from NWS by station + forecast grid
- In-memory weather cache in `weather.Service`
- Manual refresh endpoints (`POST /weather*`)
- Background refresh loop:
  - current weather every 5 minutes
  - full weather refresh every 15 minutes
- Unit normalization and wind direction mapping in `scraping/mapping.go`

## Project Layout

- `main.go`: app bootstrap, service wiring, refresh loop, and HTTP server startup
- `api/`: HTTP handlers and route registration
- `weather/`: domain types + in-memory service/cache
- `scraping/`: NWS provider, transport, and mapping logic
- `scraping/observations/`: NWS response models and JSON helpers

## Requirements

- Go `1.26` (from `go.mod`)
- Network access to `https://api.weather.gov`
- Environment variables:
  - `EMAIL` (used in NWS `User-Agent`)
  - `STATION` (example: `KNYC`)
  - `OFFICE` (example: `OKX`)
  - `GRIDX` (example: `34`)
  - `GRIDY` (example: `48`)

Optional:

- `LATLONG` (only used by experimental `ScrapeWebsite()`)

## Run

```bash
go run .
```

Server listens on `http://localhost:8080`.

## API Endpoints

- `GET /health`
- `GET /weather`
- `POST /weather`
- `GET /weather/current`
- `POST /weather/current`
- `GET /weather/hourly`
- `POST /weather/hourly`
- `GET /weather/daily`
- `POST /weather/daily`

Quick check:

```bash
curl http://localhost:8080/health
curl http://localhost:8080/weather/current
curl -X POST http://localhost:8080/weather/hourly
```

## Core Types and Data Flow

- `scraping.WeatherGovProvider` implements `weather.Provider`
- `weather.Service` owns cached `WeatherData` and refresh methods
- `api.Handler` reads from and refreshes `weather.Service`

## Development

```bash
go fmt ./...
go test ./...
```

Current test status: repository compiles and `go test ./...` passes, but there are no test files yet.

## Notes

- The service exits at startup if `GRIDX` or `GRIDY` are missing/invalid.
- Station/office/grid values must be valid NWS identifiers.
- `scraping/ScrapeWebsite()` is experimental and not part of the HTTP API path.
