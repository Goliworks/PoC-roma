package server

import (
	"fmt"
	"github.com/Goliworks/Roma/internal/config"
	"log"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
	cfg *config.Config
}

func NewServer() *Server {
	srv := new(Server)
	srv.mux = http.NewServeMux()
	srv.cfg = config.NewConfig()
	srv.mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			Handler(w, r, srv.cfg)
		})
	return srv
}

func (s *Server) Launch() {
	srv := http.Server{
		Addr:    s.cfg.Port,
		Handler: s.mux,
	}
	fmt.Printf("Launch simple server on port %v\n", s.cfg.Port)
	log.Fatal(srv.ListenAndServe())
}

func (s *Server) LaunchTLS() {
	srv := http.Server{
		Addr:    ":8443",
		Handler: s.mux,
	}
	fmt.Printf("Launch simple server on port %v\n", ":8443")
	log.Fatal(srv.ListenAndServeTLS("certs/server.crt", "certs/server.key"))
}
