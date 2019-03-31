package router

import (
	"log"
	"net/http"

	StatusHandler "github.com/hikvineh/go-api/cmd/controllers/v1/status"
	"github.com/hikvineh/go-api/pkg/types/routes"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-App-Token")
		if len(token) < 1 {
			http.Error(w, "Not Authorized!", http.StatusUnauthorized)
			return
		}

		log.Println("Inside Middleware v1")
		// log.Println(r.Header.Get("X-App-Token"))

		next.ServeHTTP(w, r)
	})
}

func GetRoutes() (SubRoute map[string]routes.SubRoutePackage) {

	/* ROUTES */
	SubRoute = map[string]routes.SubRoutePackage{
		"/v1": routes.SubRoutePackage{
			Routes: routes.Routes{
				routes.Route{"Status", "GET", "/status", StatusHandler.Index},
			},
			Middleware: Middleware,
		},
	}

	return
}
