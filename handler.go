package main

import (
	"github.com/fajarardiyanto/flt-go-router/interfaces"
	"net/http"
)

func HandlerRegencies(w http.ResponseWriter, r *http.Request) error {
	name := interfaces.GetQuery(r, "name")

	pool := NewPooling()
	villages, err := pool.GetRegencies(name)
	if err != nil {
		return err
	}

	return interfaces.JSON(w, http.StatusOK, villages)
}
