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
	srv := Server{
		mux: http.NewServeMux(),
	}
	srv.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Hello, Roma!")
	})
	return &srv
}

func (s *Server) Launch(){
	srv := http.Server{
		Addr: ":8080",
		Handler: s.mux,
	}

	fmt.Println("Launch simple server on port :8080")
	log.Fatal(srv.ListenAndServe())
}