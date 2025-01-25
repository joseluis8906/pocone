package app

import (
	"context"
	"fmt"
	stdlog "log"
	"net"
	"net/http"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joseluis8906/pocone/internal/product"
	"github.com/joseluis8906/pocone/pkg/log"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Deps struct {
	fx.In
	Config    *viper.Viper
	Log       *stdlog.Logger
	RpcServer *rpc.Server
	Products  *product.RpcService
}

func NewRpcServer() *rpc.Server {
	return rpc.NewServer()
}

func New(lc fx.Lifecycle, deps Deps) *http.Server {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", deps.Config.GetInt("http.server.port")),
		Handler: cors.Default().Handler(deps.RpcServer),
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			l, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				log.Fatalf("%s creating server bind: %v", log.Error, err)
			}

			log.Printf("%s http server listening on %s", log.Info, srv.Addr)
			go srv.Serve(l)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.Shutdown(context.Background())
			return nil
		},
	})

	return srv
}

func NewRouter() *http.ServeMux {
	return http.NewServeMux()
}
