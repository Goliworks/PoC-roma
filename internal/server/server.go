package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server{
	srv := new(Server)
	srv.mux = http.NewServeMux()
	srv.mux.HandleFunc("/", Handler)
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