package customer

import (
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joseluis8906/pocone/pkg/db"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewRpcService,
	NewRepository,
)

type (
	Deps struct {
		fx.In
		DB        *db.Database
		RpcServer *rpc.Server
	}
)

func NewRpcService(deps Deps) *RpcService {
	setDBIndexes(deps.DB.Collection(collection))

	srv := &RpcService{
		Repository: &Repository{
			db: deps.DB.Collection(collection),
		},
	}

	deps.RpcServer.RegisterName("customer", srv)
	return srv
}
