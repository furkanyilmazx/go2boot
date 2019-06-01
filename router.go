package system

import (
	"go2api/system/middleware"
	"go2api/system/router"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(apiPrefix string, routes router.Routes, middlewares middleware.MiddlewareWrapFuncs) *mux.Router {
	mainRouter := mux.NewRouter().StrictSlash(false)
	router := mainRouter.PathPrefix(apiPrefix).Subrouter()
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = middleware.HandlerMiddleware(handler, route.Middlewares)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return mainRouter
	//return middleware.RouterMiddleware(router, middlewares)
	//return handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)
}
