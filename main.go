package main

import (
	"fmt"

	"github.com/ajn2004/go-weather/scraping"
)

func main() {
	obs, _ := scraping.GetCurrentObservation("KNYC", "andrew@daleego.com")
	fmt.Printf("Observation: %v\n", obs)

}
