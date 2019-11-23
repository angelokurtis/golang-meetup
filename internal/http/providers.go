package http

import (
	"github.com/google/wire"
)

var Providers = wire.NewSet(NewHttpClient)
