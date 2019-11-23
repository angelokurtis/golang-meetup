package log

import (
	"github.com/google/wire"
)

var Providers = wire.NewSet(
	NewSugaredLogger,
	NewZapLogger,
	wire.Bind(new(Logger), new(*ZapLogger)),
)
