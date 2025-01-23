package order

import (
	myhttp "github.com/joseluis8906/pocone/pkg/http"
)

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
