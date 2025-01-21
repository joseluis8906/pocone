package order

import (
	"fmt"
	"net/http"

	"github.com/joseluis8906/pocone/pkg/db"
	myhttp "github.com/joseluis8906/pocone/pkg/http"
	"github.com/joseluis8906/pocone/pkg/log"
	"go.uber.org/fx"
)

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

func setRoutes(hf myhttp.HandlerFn) {
	middlewares := myhttp.MiddlewareList(
		myhttp.Middleware(myhttp.MiddlewarePre, myhttp.LogReq),
		myhttp.Middleware(myhttp.MiddlewarePost, myhttp.LogRes),
	)

	hf(myhttp.Route("/orders", create, middlewares...))
	hf(myhttp.Route("/orders/{id}", resume, middlewares...))
	hf(myhttp.Route("/orders/{id}/customer", addCustomer, middlewares...))
	hf(myhttp.Route("/orders/{id}/items", addItem, middlewares...))
	hf(myhttp.Route("/orders/{id}/items/{pos}", removeItem, middlewares...))
	hf(myhttp.Route("/orders/{id}/items/{pos}/quantity", changeItemQty, middlewares...))
}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var task CreateTask
	task.GenerateID()
	task.PersistOrder(r.Context())
	task.MarshalOrder()

	data, err := task.Result()
	if err != nil {
		log.Printf("%s creating order: %s", log.Error, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s\n", data)
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func addItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func removeItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func changeItemQty(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func pay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func resume(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var task GetTask
	task.ExtractID(r)
	task.SearchTheDB(r.Context())
	task.Encode()

	data, err := task.Result()
	if err != nil {
		log.Printf("%s showing order: %v", log.Error, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s\n", data)
}
