package main

import (
	"encoding/json"
	"net/http"
)

func HandlerRegencies(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	pool := NewPooling()
	villages, err := pool.GetRegencies(name)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(villages)
}
