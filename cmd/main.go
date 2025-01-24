package main

import (
	"github.com/joseluis8906/pocone/internal/app"
	"github.com/joseluis8906/pocone/internal/product"
	"github.com/joseluis8906/pocone/pkg/config"
	"github.com/joseluis8906/pocone/pkg/db"
	"github.com/joseluis8906/pocone/pkg/log"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(config.New),
		fx.Provide(log.New),
		fx.Provide(db.New),

		fx.Options(product.Module),

		fx.Provide(app.New),
		fx.Invoke(func(*app.Server) {}),
	).Run()
}
