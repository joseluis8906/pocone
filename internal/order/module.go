package order

import (
	"net/http"

	"github.com/joseluis8906/pocone/pkg/db"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewRouter)

type (
	Router struct{}

	Deps struct {
		fx.In
		Router *http.ServeMux
		DB     *db.Database
	}
)

func NewRouter(deps Deps) *Router {
	setRoutes(deps.Router.HandleFunc)
	setDBindexes(deps.DB.Collection(collection))

	return nil
}
