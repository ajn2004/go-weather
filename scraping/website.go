package scraping

import (
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

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
	if err != nil {
		return "", fmt.Errorf("error making website request: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected error code fetching website: %d", res.StatusCode)
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
