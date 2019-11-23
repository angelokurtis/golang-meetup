//+build wireinject

package main

import (
	"github.com/angelokurtis/golang-meetup/internal/http"
	"github.com/angelokurtis/golang-meetup/internal/log"
	"github.com/angelokurtis/golang-meetup/pkg/location"
	"github.com/angelokurtis/golang-meetup/pkg/weather"
	"github.com/google/wire"
)

func Initialize() weather.Checker {
	wire.Build(http.Providers, log.Providers, weather.Providers, location.Providers)
	return nil
}
