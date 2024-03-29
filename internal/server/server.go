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
	mux := generateMux(Handler, s.cfg)
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
	mux := generateMux(HandlerTLS, s.cfg)
	srv := http.Server{
		Addr:      s.cfg.PortTLS,
		Handler:   mux,
		TLSConfig: s.cfg.TLSConf,
	}
	fmt.Printf("Launch HTTPS server on port %v\n", s.cfg.PortTLS)
	log.Fatal(srv.ListenAndServeTLS("", ""))
}

func generateMux(h HandlerDefinition, cfg *config.Config) (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			h(w, r, cfg)
		})
	return
}
