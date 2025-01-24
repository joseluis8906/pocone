package order

import (
	"github.com/joseluis8906/pocone/pkg/db"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewRouter)

type (
	Router struct{}

	Deps struct {
		fx.In
		DB *db.Database
	}
)

func NewRouter(deps Deps) *Router {
	setDBindexes(deps.DB.Collection(collection))

	return nil
}
