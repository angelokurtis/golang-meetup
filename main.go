package main

import (
	"log"
)

func main() {
	weatherChecker := Initialize() // initializing by injectors
	weather, err := weatherChecker.CheckByCurrentLocation()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("o clima atual é de %.2f°C e %s", weather.Temp, weather.Description)
}
