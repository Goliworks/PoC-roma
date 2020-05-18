package server

import (
	"fmt"
	"net/http"
)

func badGateway(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Println(err)
	fmt.Println(r.URL.Path)
	w.WriteHeader(http.StatusBadGateway)
	fmt.Fprint(w,
		"<div style='text-align:center; margin-top:100px;"+
			"font-family:helvetica, arial;'>"+
			"<div><h1>Error 502</h1>"+
			"<span>Bad gateway</span></div></div>")
}
