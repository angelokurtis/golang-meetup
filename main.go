package main

import (
	"fmt"
	"log"
)

func main() {
	app := Initialize()
	logger := app.Logger()
	forecastChecker := app.ForecastChecker()
	locationFinder := app.LocationFinder()

	l, err := locationFinder.WhereAmI()
	if err != nil {
		log.Fatal(err)
	}
	w, err := forecastChecker.CheckByGeoCoordinates(l.Latitude, l.Longitude)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info(fmt.Sprintf("o clima atual é de %.2f°C com %s", w.Temp, w.Description))
}
