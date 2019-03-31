package app

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/gorilla/handlers"
	"github.com/hikvineh/go-api/cmd/system/router"
)

// Server struct
type Server struct {
	port string
	Db   *xorm.Engine
}

// NewServer all vals
func NewServer() Server {
	return Server{}
}

// Init all vals
func (s *Server) Init(port string, db *xorm.Engine) {
	log.Println("Initializing Server...")
	s.port = ":" + port
	s.Db = db
}

// Start all vals
func (s *Server) Start() {
	log.Println("Starting Server on port" + s.port)

	r := router.NewRouter()

	r.Init()

	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Cache-Control", "X-App-Token"}),
		handlers.ExposedHeaders([]string{""}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(r.Router))
	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

	newServer := &http.Server{
		Handler:      handler,
		Addr:         "0.0.0.0" + s.port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(newServer.ListenAndServe())
	// http.ListenAndServe(s.port, handler)

}
