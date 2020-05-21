package server

import (
	"fmt"
	"github.com/Goliworks/Roma/internal/config"
	"net/http"
	"net/url"
)

type HandlerDefinition func(w http.ResponseWriter, r *http.Request, cfg *config.Config)

func Handler(w http.ResponseWriter, r *http.Request, cfg *config.Config) {
	if cfg.AutoTLS {
		u, _ := url.Parse(fmt.Sprintf("http://%v", r.Host))
		port := ""
		if cfg.PortTLS != config.DefaultTLSPort {
			port = cfg.PortTLS
		}
		red := fmt.Sprintf("https://%v%v%v", u.Hostname(), port, r.RequestURI)
		http.Redirect(w, r, red, http.StatusMovedPermanently)
		return
	}
	serveRProxy(w, r, cfg)
}

func HandlerTLS(w http.ResponseWriter, r *http.Request, cfg *config.Config) {
	// rules for TLS
	serveRProxy(w, r, cfg)
}
