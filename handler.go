package main

import (
	"github.com/fajarardiyanto/flt-go-router/interfaces"
	"net/http"
)

type handler struct {
	repo ApiIndonesiaArea
}

func NewHandler(repo ApiIndonesiaArea) *handler {
	return &handler{repo}
}

func (h *handler) HandlerProvinces(w http.ResponseWriter, r *http.Request) error {
	name := interfaces.GetQuery(r, "name")

	prov, err := h.repo.GetProvinces(name)
	if err != nil {
		return err
	}

	return interfaces.JSON(w, http.StatusOK, prov)
}

func (h *handler) HandlerRegencies(w http.ResponseWriter, r *http.Request) error {
	name := interfaces.GetQuery(r, "name")

	villages, err := h.repo.GetRegencies(name)
	if err != nil {
		return err
	}

	return interfaces.JSON(w, http.StatusOK, villages)
}
