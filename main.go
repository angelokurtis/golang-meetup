package main

import (
	"fmt"
	"github.com/angelokurtis/golang-meetup/pkg/weather"
	"log"
)

func main() {
	weatherChecker := weather.NewOpenWeather()
	w, err := weatherChecker.CheckByCurrentLocation()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("o clima atual é de %.2f°C e %s\n", w.Temp, w.Description)
}
