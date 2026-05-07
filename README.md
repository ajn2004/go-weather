# go-weather

`go-weather` is a Go weather client focused on pulling NOAA/NWS weather data from `api.weather.gov` and mapping it into app-friendly domain types.

## Features

- Fetch latest station observation data by station ID
- Fetch hourly forecast data by forecast office + grid coordinates
- Normalize common units (temperature, wind speed, pressure, visibility)
- Map wind direction into typed cardinal values
- Support mixed API payload shapes (for example, temperature and wind speed variants)

## Project Layout

- `main.go`: Example entry point showing current observation and hourly forecast calls
- `scraping/`: Data fetching + mapping layer
- `scraping/observations/`: API response models and custom JSON unmarshal helpers
- `weather/`: App/domain weather types
- `api/`: Placeholder package for future API layer

## Requirements

- Go `1.26` (as defined in `go.mod`)
- Network access to `https://api.weather.gov`
- A contact email set as `EMAIL` (used in the request `User-Agent`)

## Setup

```bash
git clone https://github.com/ajn2004/go-weather.git
cd go-weather
go mod download
```

Set your environment variable:

```bash
export EMAIL="you@example.com"
```

Optional (for the website scraping helper):

```bash
export LATLONG="lat=40.7128&lon=-74.0060"
```

## Run

Run the example program:

```bash
go run .
```

The example currently:

- Gets latest observation for station `KNYC`
- Gets hourly forecast for office `OKX` at grid `34,48`
- Prints one current observation and the first hourly forecast period

## Usage Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/ajn2004/go-weather/scraping"
)

func main() {
	obs, err := scraping.GetCurrentObservation("KNYC")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current: %+v\n", obs)

	hourly, err := scraping.GetHourlyForecast("OKX", 34, 48)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Next hour: %+v\n", hourly[0])
}
```

## Public Functions

- `scraping.GetCurrentObservation(stationID string) (weather.CurrentWeather, error)`
- `scraping.GetHourlyForecast(forecastOfficeId string, forecastGridX int, forecastGridY int) ([]weather.HourlyForecast, error)`
- `scraping.ScrapeWebsite() (string, error)` (experimental helper)

## Notes and Limitations

- `main.go` currently ignores returned errors; production callers should always check errors.
- `ScrapeWebsite()` is exploratory and not a stable API.
- There are currently no automated tests in this repository.
- Station IDs and forecast grid coordinates must be valid NWS identifiers.

## Development

Common commands:

```bash
go fmt ./...
go test ./...
```

If you add tests or new packages, prefer keeping parsing, transport, and domain mapping responsibilities separated to preserve maintainability.
