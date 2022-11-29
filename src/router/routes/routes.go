package routes

import (
	"net/http"

	"github.com/margen2/shorgot/src/middlewares"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

// Config adds each route to the router and adds the Logger  and authentication middleware to them
func Config(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, linkRoutes...)

	for _, route := range routes {
		if route.RequiresAuthentication {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}
	return r
}
