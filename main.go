package main

import (
	"flag"
	"github.com/gorilla/mux"
	"net/http"
)

var bindAddr string

func main() {
	route := mux.NewRouter()
	flag.StringVar(&bindAddr, "bind", ":8081", "bind addr")

	route.HandleFunc("/regencies", SetMiddlewareJSON(HandlerRegencies))

	GetLogger().Info("SERVER STARTING AT %s", bindAddr)
	http.ListenAndServe(bindAddr, route)
}
