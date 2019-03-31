package router

import (
	"net/http"

	HomeHandler "github.com/hikvineh/go-api/cmd/controllers/home"
	"github.com/hikvineh/go-api/pkg/types/routes"
)

func MiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func GetRoutes() routes.Routes {

	return routes.Routes{
		routes.Route{"Home", "GET", "/", HomeHandler.Index},
	}
}
