package main

import (
	"flag"
	"github.com/fajarardiyanto/flt-go-router/lib"
)

var bindAddr string

func main() {
	flag.StringVar(&bindAddr, "bind", ":8081", "bind addr")
	//bindAddr := os.Getenv("PORT")

	router := lib.New("v1.0.0")
	router.Use(MiddlewareLogger(), MiddlewareError())

	svc := NewService()
	h := NewHandler(svc)

	router.GET("", h.HandlerInfo)
	router.GET("/provinces", h.HandlerProvinces)
	router.GET("/regencies", h.HandlerRegencies)
	router.GET("/districts", h.HandlerDistricts)
	router.GET("/villages", h.HandlerVillages)

	router.Run(bindAddr)
}
