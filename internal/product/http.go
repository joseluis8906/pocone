package product

import (
	myhttp "github.com/joseluis8906/pocone/pkg/http"
)

func setRoutes(hf myhttp.HandlerFn) {
	middlewares := myhttp.MiddlewareList(
		myhttp.Middleware(myhttp.MiddlewarePre, myhttp.LogReq),
		myhttp.Middleware(myhttp.MiddlewarePost, myhttp.LogRes),
	)

	hf(myhttp.Route("/products", all, middlewares...))
}
