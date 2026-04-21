package scraping

import "github.com/ajn2004/go-weather/weather"

// MapToWeatherData maps the scraped data to the WeatherData struct
func mapObservationToCurrentWeather(resp ObservationResponse) weather.CurrentWeather {
	var cw weather.CurrentWeather
	if resp.Properties.Temperature.Value != nil {
		// check if units are in celsius, if not convert from fahrenheit to celsius
		cw.Temperature = enforceCelsius(*resp.Properties.Temperature.Value, resp.Properties.Temperature.Unit)

	}
	if resp.Properties.Humidity.Value != nil {
		cw.Humidity = *resp.Properties.Humidity.Value
	}
	if resp.Properties.WindSpeed.Value != nil {
		// check if units are in m/s, if not convert into m/s
		cw.WindSpeed = enforceMetersPerSec(*resp.Properties.WindSpeed.Value, resp.Properties.WindSpeed.Unit)
	}
	if resp.Properties.Pressure.Value != nil {
		// check if units are in hPa, if not convert from Pa to hPa
		// we want pressure in atm, so convert from hPa or Pa to atm
		cw.Pressure = enforceAtmospheres(*resp.Properties.Pressure.Value, resp.Properties.Pressure.Unit)
	}
	if resp.Properties.DewPoint.Value != nil {
		cw.DewPoint = enforceCelsius(*resp.Properties.DewPoint.Value, resp.Properties.DewPoint.Unit)
	}
	if resp.Properties.Visibility.Value != nil {
		// check if units are in km, if not convert from m to km
		// we want visibility in km, so convert from m or mi to km
		cw.Visibility = enforceKilometers(*resp.Properties.Visibility.Value, resp.Properties.Visibility.Unit)
	}

	cw.Condition = resp.Properties.Condition
	cw.LastUpdate = resp.Properties.Timestamp

	// Map wind direction from degrees to cardinal direction
	if resp.Properties.WindDirection.Value != nil {
		cw.WindDirection = degreesToCardinal(*resp.Properties.WindDirection.Value)
	} else {
		cw.WindDirection = weather.NA // default value if wind direction is missing
	}

	return cw
}

func enforceCelsius(temp float64, unit string) float64 {
	if unit == "unit:degC" {
		return temp
	}
	if unit == "unit:degF" {
		return (temp - 32) * 5.0 / 9.0
	}
	return temp // if unit is unrecognized, return the original value
}

func enforceMetersPerSec(speed float64, unit string) float64 {
	if unit == "unit:km_h-1" {
		return speed / 3.6
	}
	if unit == "unit:mph" {
		return speed * 1.60934 / 3.6
	}
	return speed // if unit is unrecognized, return the original value
}

func enforceAtmospheres(pressure float64, unit string) float64 {
	if unit == "unit:hPa" {
		return pressure / 1013.25
	}
	if unit == "unit:Pa" {
		return pressure / 101325.0
	}
	return pressure // if unit is unrecognized, return the original value
}

func enforceKilometers(distance float64, unit string) float64 {
	if unit == "unit:km" {
		return distance
	}
	if unit == "unit:m" {
		return distance / 1000.0
	}
	if unit == "unit:mi" {
		return distance * 1.60934
	}
	return distance // if unit is unrecognized, return the original value
}

func degreesToCardinal(degrees float64) weather.WindDirection {
	switch {
	case degrees >= 337.5 || degrees < 22.5:
		return weather.North
	case degrees < 67.5:
		return weather.Northeast
	case degrees < 112.5:
		return weather.East
	case degrees < 157.5:
		return weather.Southeast
	case degrees < 202.5:
		return weather.South
	case degrees < 247.5:
		return weather.Southwest
	case degrees < 292.5:
		return weather.West
	case degrees < 337.5:
		return weather.Northwest
	default:
		return weather.NA
	}
}
