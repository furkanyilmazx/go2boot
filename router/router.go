package router

import (
	"go2api/system/middleware"
	"log"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Middlewares middleware.Middlewares
}

type Routes []Route

func ConcatRoutes(routes ...Routes) Routes {
	var aa Routes = Routes{}
	for _, i := range routes {
		aa = append(aa, i...)
		for _, route := range i {
			log.Println(route.Method, route.Pattern)
		}
	}
	return aa
}
