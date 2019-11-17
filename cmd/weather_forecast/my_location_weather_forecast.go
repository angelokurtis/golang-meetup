package main

import (
	"fmt"
	"github.com/angelokurtis/golang-meetup/internal/log"
	"github.com/angelokurtis/golang-meetup/internal/temperatures"
	"github.com/angelokurtis/golang-meetup/pkg/forecast"
	"github.com/angelokurtis/golang-meetup/pkg/location"
)

type MyLocationWeatherForecast struct {
	log             *log.Logger
	locationFinder  *location.Finder
	forecastChecker *forecast.Checker
}

func NewMyLocationWeatherForecast(l *log.Logger, lf *location.Finder, wf *forecast.Checker) *MyLocationWeatherForecast {
	return &MyLocationWeatherForecast{log: l, locationFinder: lf, forecastChecker: wf}
}

func (wf *MyLocationWeatherForecast) ForTomorrow() (string, error) {
	l, err := (*wf.locationFinder).WhereAmI()
	if err != nil {
		return "", err
	}
	f, err := (*wf.forecastChecker).CheckByGeoCoordinates(l.Latitude, l.Longitude)
	if err != nil {
		return "", err
	}
	min := temperatures.Fahrenheit(f.Temp.AverageMin)
	max := temperatures.Fahrenheit(f.Temp.AverageMax)
	return fmt.Sprintf("min %f | max %f", min.ToCelsius(), max.ToCelsius()), nil
}
