package scraping

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/ajn2004/go-weather/scraping/observations"
	"github.com/ajn2004/go-weather/weather"
)

func GetCurrentObservation(stationID string) (weather.CurrentWeather, error) {
	var obs weather.CurrentWeather

	// Define the base URL for the API endpoint, inserting the station ID
	baseURL := "https://api.weather.gov/stations/" + stationID + "/observations/latest"
	// Create an HTTP client with a timeout to avoid hanging requests
	client := getClient()
	req, err := getRequest(baseURL)
	if err != nil {
		return obs, err
	}

	// Make the request and check for errors
	res, err := client.Do(req)
	if res.StatusCode != http.StatusOK {
		return obs, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	defer res.Body.Close() // ensure the response body is closed after we're done with it

	// Decode the JSON response into the ObservationResponse struct
	var resp observations.ObservationResponse

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return obs, fmt.Errorf("error decoding response: %w", err)
	}

	fmt.Println("Date:", res.Header.Get("Date"))
	fmt.Println("Age:", res.Header.Get("Age"))
	fmt.Println("Cache-Control:", res.Header.Get("Cache-Control"))
	fmt.Println("Last-Modified:", res.Header.Get("Last-Modified"))
	fmt.Println("Expires:", res.Header.Get("Expires"))
	obs = mapObservationToCurrentWeather(resp)

	return obs, nil
}

func GetHourlyForecast(forecastOfficeId string, forecastGridX int, forecastGridY int) ([]weather.HourlyForecast, error) {
	var forecast []weather.HourlyForecast

	url := fmt.Sprintf("https://api.weather.gov/gridpoints/%s/%d,%d/forecast/hourly", forecastOfficeId, forecastGridX, forecastGridY)
	client := getClient()

	req, err := getRequest(url)
	if err != nil {
		return forecast, err
	}

	res, err := client.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return forecast, fmt.Errorf("error fetching hourly forecast: %w", err)
	}
	defer res.Body.Close()

	var resp observations.HourlyForecastResponse

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return forecast, fmt.Errorf("error decoding hourly forecast response: %w", err)
	}

	forecast = mapHourlyResponseToHourlyForecast(resp, 12)
	return forecast, nil
}

// daily scraping
func GetDailyForecast(forecastOfficeId string, forecastGridX int, forecastGridY int) ([]weather.DailyForecast, error) {
	var forecast []weather.DailyForecast

	url := fmt.Sprintf("https://api.weather.gov/gridpoints/%s/%d,%d/forecast", forecastOfficeId, forecastGridX, forecastGridY)
	client := getClient()

	req, err := getRequest(url)
	if err != nil {
		return forecast, err
	}

	res, err := client.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return forecast, fmt.Errorf("error fetching daily forecast: %w", err)
	}
	defer res.Body.Close()

	var resp observations.DailyForecastResponse
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return forecast, fmt.Errorf("error decoding daily forecast response: %w", err)
	}

	forecast = mapDailyResponseToDailyForecast(resp, 7)
	return forecast, nil
}

func getClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

func getRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil,
			fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("User-Agent", "go-weather-scraper/1.0 (contact:"+os.Getenv("EMAIL")+")")
	req.Header.Set("Accept", "application/geo+json")
	return req, nil
}

func ScrapeWebsite() (string, error) {
	// This is to scrape weather.gov directly, however it seems apparent that it grabs the same information from the same source as the API, so it may not be necessary to implement this function. However, if there are any additional data points that we want to scrape from the website that are not available in the API, we can implement this function to do so.
	// This function can be implemented to scrape data from a website if needed
	baseURL := "https://forecast.weather.gov/"
	client := getClient()
	req, err := getRequest(baseURL + "MapClick.php?" + os.Getenv("LATLONG"))
	if err != nil {
		return "", err
	}

	res, err := client.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error fetching website: %w", err)
	}
	defer res.Body.Close()

	// Here you would add the logic to parse the HTML response and extract the desired data
	// print the response text for now
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", fmt.Errorf("error parsing HTML: %w", err)
	}

	// Example: Extract the current temperature from the page
	fmt.Println("Scraping website for current temperature...")
	currentTemp := doc.Find(".myforecast-current-lrg").Text()
	fmt.Printf("Current Temperature: %s\n", currentTemp)
	// currentHumidity := doc.Find("#current_conditions_detail").Find("table").Find("tr").Eq(0).Find("td").Text()
	// fmt.Printf("Current Humidity: %s\n", BcurrentHumidity)
	return "", nil
}
