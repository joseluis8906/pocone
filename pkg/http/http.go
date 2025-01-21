package http

import (
	"net/http"
)

type (
	HandlerFn = func(string, func(http.ResponseWriter, *http.Request))
	Handler   = func(w http.ResponseWriter, r *http.Request)

	middlewareType int8
	middleware     struct {
		kind middlewareType
		fn   Handler
	}
)

func MiddlewareList(middlewares ...middleware) []middleware {
	return middlewares
}

func Middleware(t middlewareType, h Handler) middleware {
	return middleware{kind: t, fn: h}
}

const (
	MiddlewarePre  middlewareType = 0
	MiddlewarePost middlewareType = 1
)

func Route(path string, h Handler, middlewares ...middleware) (string, Handler) {
	return path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := newResponseWriter(w)

		mdsByType := map[middlewareType][]middleware{}
		for _, md := range middlewares {
			mdsByType[md.kind] = append(mdsByType[md.kind], md)
		}

		mds := mdsByType[MiddlewarePre]
		for _, next := range mds {
			next.fn(rw, r)
		}

		h(rw, r)

		mds = mdsByType[MiddlewarePost]
		for _, next := range mds {
			next.fn(rw, r)
		}
	})
}
