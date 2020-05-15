package server

import (
	"fmt"
	"github.com/Goliworks/Roma/internal/config"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Handler (w http.ResponseWriter, r *http.Request, cfg *config.Config) {
	ptc := "http"
	inUrl, _ := url.Parse(fmt.Sprintf("%v://%v", ptc, r.Host))
	destUrl, _ := url.Parse(fmt.Sprintf("http://%v", cfg.Destinations[inUrl.Hostname()]))

	fmt.Printf("incoming host : %v\n", r.Host)
	fmt.Printf("destination host : %v\n", destUrl.Host)

	proxy := httputil.NewSingleHostReverseProxy(destUrl)
	proxy.ServeHTTP(w, r)
}