package product

import (
	"github.com/joseluis8906/pocone/pkg/db"
	"go.uber.org/fx"
)

var Module = fx.Provide(New)

type (
	Deps struct {
		fx.In
		DB *db.Database
	}
)

func New(deps Deps) *RpcService {
	setDBIndexes(deps.DB.Collection(collection))

	return &RpcService{Repository: &Repository{db: deps.DB.Collection(collection)}}
}
