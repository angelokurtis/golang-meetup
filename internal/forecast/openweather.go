package forecast

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/angelokurtis/golang-meetup/internal/log"
	"github.com/angelokurtis/golang-meetup/pkg/forecast"
	"net/http"
	"os"
)

type OpenWeather struct {
	log *log.Logger
}

func (o *OpenWeather) CheckByGeoCoordinates(latitude float64, longitude float64) (*forecast.Forecast, error) {
	(*o.log).Info(fmt.Sprintf("checking weather for latitude %.2f and longitude %.2f", latitude, longitude))

	ak := os.Getenv("OPEN_WEATHER_API_ACCESS_KEY")
	r, err := http.Get(fmt.Sprintf("https://pro.openweathermap.org/data/2.5/climate/month?lat=%f&lon=%f&lang=pt&appid=%s", latitude, longitude, ak))
	if err != nil {
		return nil, err
	}
	b := r.Body
	defer b.Close()
	f := &Forecasts{}
	err = json.NewDecoder(b).Decode(f)
	if err != nil {
		return nil, err
	}
	forecasts := f.List
	if len(forecasts) < 1 {
		return nil, errors.New("forecasts is empty")
	}
	fct := forecasts[0]
	(*o.log).Info(fmt.Sprintf("forecast found for %d", fct.Dt))
	return &fct, nil
}

type Forecasts struct {
	Cod  string `json:"cod"`
	City struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Coord struct {
			Lon float64 `json:"lon"`
			Lat float64 `json:"lat"`
		} `json:"coord"`
		Country string `json:"country"`
	} `json:"city"`
	Message float64             `json:"message"`
	List    []forecast.Forecast `json:"list"`
}
