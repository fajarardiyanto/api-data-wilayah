package main

import (
	"flag"
	"github.com/gorilla/mux"
	"net/http"
	"runtime"
	"time"
)

var bindAddr string

func main() {
	go func() {
		for {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)

			GetLogger().Info("Alloc = %v MiB (%v KiB)", bToMb(m.Alloc), bToKb(m.Alloc))
			GetLogger().Info("TotalAlloc = %v MiB (%v KiB)", bToMb(m.TotalAlloc), bToKb(m.TotalAlloc))
			GetLogger().Info("Sys = %v MiB (%v KiB)", bToMb(m.Sys), bToKb(m.Sys))
			GetLogger().Info("NumGC = %v\n", m.NumGC)

			time.Sleep(1 * time.Second)
		}
	}()

	route := mux.NewRouter()
	flag.StringVar(&bindAddr, "bind", ":8081", "bind addr")

	route.HandleFunc("/regencies", SetMiddlewareJSON(HandlerRegencies))

	GetLogger().Info("SERVER STARTING AT %s", bindAddr)
	http.ListenAndServe(bindAddr, route)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func bToKb(b uint64) uint64 {
	return b / 1024
}
