package server

import (
	"fmt"
	"github.com/Goliworks/Roma/internal/config"
	"log"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server{
	srv := new(Server)
	srv.mux = http.NewServeMux()
	cfg := config.NewConfig()
	srv.mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		Handler(w, r, cfg)
	})
	return srv
}

func (s *Server) Launch(){
	srv := http.Server{
		Addr: ":8080",
		Handler: s.mux,
	}
	fmt.Println("Launch simple server on port :8080")
	log.Fatal(srv.ListenAndServe())
}