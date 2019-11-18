package forecast

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/angelokurtis/golang-meetup/internal/log"
	"net/http"
	"os"
)

type OpenWeather struct {
	log log.Logger
}

func NewOpenWeather(log log.Logger) *OpenWeather {
	return &OpenWeather{log: log}
}

func (o *OpenWeather) CheckByGeoCoordinates(latitude float64, longitude float64) (*Weather, error) {
	o.log.Info(fmt.Sprintf("verificando clima atual nas coordenadas de latitude %.2f e longitude %.2f", latitude, longitude))

	ak := os.Getenv("OPEN_WEATHER_API_ACCESS_KEY")
	r, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&APPID=%s&units=metric&lang=pt", latitude, longitude, ak))
	if err != nil {
		return nil, err
	}
	if r.StatusCode == http.StatusOK {
		b := r.Body
		defer b.Close()
		w := &OpenWeatherResponse{}
		err = json.NewDecoder(b).Decode(w)
		if err != nil {
			return nil, err
		}
		weathers := w.Weather
		if len(weathers) < 1 {
			return nil, errors.New("forecasts is empty")
		}
		weather := weathers[0]
		o.log.Info("clima atual encontrado")
		return &Weather{
			Description: weather.Description,
			Temp:        w.Main.Temp,
		}, nil
	}
	return nil, errors.New(fmt.Sprintf("OpenWeather responded unexpectedly (HTTP code %d)", r.StatusCode))
}

type OpenWeatherResponse struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}
