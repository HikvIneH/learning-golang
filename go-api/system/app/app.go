package app

import (
	"log"
	"net/http"

	"github.com/go-xorm/xorm"
	"github.com/gorilla/mux"
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

// HomeHandler to handle router /
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

// Start all vals
func (s *Server) Start() {
	log.Println("Starting Server on port" + s.port)

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)

	http.ListenAndServe(s.port, router)

}
