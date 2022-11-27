package router

import (
	"github.com/margen2/shorgot/api/src/router/routes"

	"github.com/gorilla/mux"
)

// GenerateRouter generates a router
func GenerateRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
