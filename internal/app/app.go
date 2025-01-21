package app

import (
	"context"
	"fmt"
	stdlog "log"
	"net"
	"net/http"

	"github.com/joseluis8906/pocone/internal/order"
	"github.com/joseluis8906/pocone/internal/product"
	"github.com/joseluis8906/pocone/pkg/log"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Deps struct {
	fx.In
	Config        *viper.Viper
	Router        *http.ServeMux
	Log           *stdlog.Logger
	OrderRouter   *order.Router
	ProductRouter *product.Router
}

func New(lc fx.Lifecycle, deps Deps) *http.Server {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", deps.Config.GetInt("http.server.port")),
		Handler: deps.Router,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				log.Fatalf("%s creating server bind: %v", log.Error, err)
			}

			log.Printf("%s http server listening on: %s", log.Info, srv.Addr)
			go srv.Serve(ln)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.Shutdown(ctx)

			return nil
		},
	})

	return srv
}

func NewRouter() *http.ServeMux {
	return http.NewServeMux()
}
