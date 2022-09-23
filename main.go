package main

import (
	"flag"
	"github.com/fajarardiyanto/flt-go-router/lib"
)

var bindAddr string

func main() {
	flag.StringVar(&bindAddr, "bind", ":8081", "bind addr")

	router := lib.New("v1.0.0")
	router.Use(MiddlewareLogger(), MiddlewareError())

	router.GET("/regencies", HandlerRegencies)

	router.Run(bindAddr)
}
