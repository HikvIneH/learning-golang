package router

import (
	"github.com/gorilla/mux"
	V1SubRoutes "github.com/hikvineh/go-api/cmd/controllers/v1/router"
	"github.com/hikvineh/go-api/pkg/types/routes"
)

// Router Define Struct
type Router struct {
	Router *mux.Router
}

// Init function
func (r *Router) Init() {
	r.Router.Use(MiddleWare)

	baseRoutes := GetRoutes()
	for _, route := range baseRoutes {
		r.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	V1SubRoutes := V1SubRoutes.GetRoutes()
	for name, pack := range V1SubRoutes {
		r.AttachSubRouterWithMiddleWare(name, pack.Routes, pack.Middleware)
	}
}

// AttachSubRouterWithMiddleWare manage subroutes
// SubRouter function from gorilla mux
func (r *Router) AttachSubRouterWithMiddleWare(path string, subroutes routes.Routes, middleware mux.MiddlewareFunc) (SubRouter *mux.Router) {

	SubRouter = r.Router.PathPrefix(path).SubRouter()
	SubRouter.Use(Middleware)

	for _, route := range subroutes {
		SubRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
}

// NewRouter /
func NewRouter() (r Router) {
	r.Router = mux.NewRouter().StrictSlash(true)

	return
}
