package forecast

import (
	"github.com/google/wire"
)

var Providers = wire.NewSet(NewOpenWeather, wire.Bind(new(Checker), new(*OpenWeather)))
