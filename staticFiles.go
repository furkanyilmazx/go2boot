package system

import (
	"go2api/system/config"
	"net/http"

	"github.com/gorilla/mux"
)

func ServeStaticFiles(router *mux.Router, config *config.Go2Config) {
	for _, statics := range config.StaticFiles {
		for uri, location := range statics.Locations {
			pathPrefix := "/" + uri + "/"
			router.PathPrefix(pathPrefix).Handler(http.StripPrefix(pathPrefix,
				http.FileServer(http.Dir(statics.Root+location))))
		}
	}
}
