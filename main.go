package main

import (
	"github.com/angelokurtis/golang-meetup/pkg/weather"
	"log"
)

func main() {
	weatherChecker := weather.NewOpenWeather()
	w, err := weatherChecker.CheckByCurrentLocation()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("o clima atual é de %.2f°C e %s", w.Temp, w.Description)
}
