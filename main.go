package main

import (
	"fmt"
	"github.com/Goliworks/Roma/internal/server"
)

func main(){
	fmt.Println()
	srv := server.NewServer()
	srv.Launch()
}