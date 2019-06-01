package middleware

import (
	"net/http"

	"github.com/gorilla/handlers"
)

type MiddlewareFunc func(http.ResponseWriter, *http.Request)
type MiddlewareFuncs []MiddlewareFunc
type MiddlewareWrapFunc func(http.Handler) http.Handler
type MiddlewareWrapFuncs []MiddlewareWrapFunc

type Middlewares struct {
	WrapReq   MiddlewareWrapFuncs
	BeforeReq MiddlewareFuncs
	AfterReq  MiddlewareFuncs
}

func HandlerMiddleware(h http.Handler, middlewares Middlewares) http.Handler {
	var temp http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, m := range middlewares.BeforeReq {
			m(w, r)
		}
		h.ServeHTTP(w, r)
		for _, m := range middlewares.AfterReq {
			m(w, r)
		}
	})
	for _, m := range middlewares.WrapReq {
		temp = m(temp)
	}
	return temp
}

func RouterMiddleware(h http.Handler, middlewares MiddlewareWrapFuncs) http.Handler {
	for _, m := range middlewares {
		h = m(h)
	}

	return handlers.RecoveryHandler()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}))
}
