package location

import (
	"github.com/google/wire"
)

var Providers = wire.NewSet(NewIPStack, wire.Bind(new(Finder), new(*IPStack)))
