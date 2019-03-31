package app

import (
	"log"
	"net/http"

	"github.com/hikvineh/go-api/cmd/system/router"

	"github.com/go-xorm/xorm"
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

	http.ListenAndServe(s.port, r.Router)

}
