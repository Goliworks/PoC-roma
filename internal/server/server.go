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
	srv := http.Server{
		Addr:      s.cfg.PortTLS,
		Handler:   s.mux,
		TLSConfig: s.cfg.TLSConf,
	}
	fmt.Printf("Launch HTTPS server on port %v\n", s.cfg.PortTLS)
	log.Fatal(srv.ListenAndServeTLS("", ""))
}
