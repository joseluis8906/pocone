package order

import (
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joseluis8906/pocone/pkg/db"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewRpcService,
	NewRepository,
)

type (
	Deps struct {
		fx.In
		Conf      *viper.Viper
		DB        *db.Database
		RpcServer *rpc.Server
	}
)

func NewRpcService(deps Deps) *RpcService {
	setDBindexes(deps.DB.Collection(collection))

	srv := &RpcService{
		OrderRepo: &Repository{
			db: deps.DB.Collection(collection),
		},
	}

	deps.RpcServer.RegisterName("order", srv)
	return srv
}
