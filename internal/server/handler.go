package server

import (
	"github.com/Goliworks/Roma/internal/config"
	"net/http"
)

type HandlerDefinition func(w http.ResponseWriter, r *http.Request, cfg *config.Config)

func Handler(w http.ResponseWriter, r *http.Request, cfg *config.Config) {
	// rules for non TLS
	serveRProxy(w, r, cfg)
}

func HandlerTLS(w http.ResponseWriter, r *http.Request, cfg *config.Config) {
	// rules for TLS
	serveRProxy(w, r, cfg)
}
