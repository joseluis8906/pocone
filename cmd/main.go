package main

import (
	"net/http"

	"github.com/joseluis8906/pocone/internal/app"
	"github.com/joseluis8906/pocone/internal/customer"
	"github.com/joseluis8906/pocone/internal/order"
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

		fx.Options(customer.Module),
		fx.Options(product.Module),
		fx.Options(order.Module),

		fx.Provide(app.NewRpcServer),
		fx.Provide(app.New),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
