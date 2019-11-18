//+build wireinject

package main

import (
	"github.com/angelokurtis/golang-meetup/internal/log"
	"github.com/angelokurtis/golang-meetup/pkg/forecast"
	"github.com/angelokurtis/golang-meetup/pkg/location"
	"github.com/google/wire"
)

func Initialize() *App {
	wire.Build(log.Providers, forecast.Providers, location.Providers, NewApp)
	return &App{}
}

type App struct {
	logger          log.Logger
	locationFinder  location.Finder
	forecastChecker forecast.Checker
}

func NewApp(logger log.Logger, locationFinder location.Finder, forecastChecker forecast.Checker) *App {
	return &App{logger: logger, locationFinder: locationFinder, forecastChecker: forecastChecker}
}

func (a *App) Logger() log.Logger {
	return a.logger
}

func (a *App) ForecastChecker() forecast.Checker {
	return a.forecastChecker
}

func (a *App) LocationFinder() location.Finder {
	return a.locationFinder
}
