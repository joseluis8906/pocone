package order

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(NewRouter)
