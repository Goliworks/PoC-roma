package main

import (
	"github.com/Goliworks/Roma/internal/server"
)

func main() {
	srv := server.NewServer()
	srv.Launch()
}
