package server

import (
	"fmt"
	"github.com/Goliworks/Roma/internal/config"
	"log"
	"net/http"
)

type Server struct {
	cfg *config.Config
}

func NewServer() *Server {
	srv := new(Server)
	srv.cfg = config.NewConfig()
	return srv
}

func (s *Server) Launch() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			Handler(w, r, s.cfg)
		})
	srv := http.Server{
		Addr:    s.cfg.Port,
		Handler: mux,
	}
	fmt.Printf("Launch server on port %v\n", s.cfg.Port)
	log.Fatal(srv.ListenAndServe())
}

func (s *Server) LaunchTLS() {
	if len(s.cfg.TLSConf.Certificates) == 0 {
		fmt.Printf(
			"No TLS certificate has been registred.\n" +
				"Server cannot run HTTPS mode.\n")
		return
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			HandlerTLS(w, r, s.cfg)
		})
	srv := http.Server{
		Addr:      s.cfg.PortTLS,
		Handler:   mux,
		TLSConfig: s.cfg.TLSConf,
	}
	fmt.Printf("Launch HTTPS server on port %v\n", s.cfg.PortTLS)
	log.Fatal(srv.ListenAndServeTLS("", ""))
}
