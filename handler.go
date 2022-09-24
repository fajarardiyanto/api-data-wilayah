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

	return interfaces.JSON(w, http.StatusOK, Response{
		Total: len(prov),
		Data:  prov,
	})
}

func (h *handler) HandlerRegencies(w http.ResponseWriter, r *http.Request) error {
	name := interfaces.GetQuery(r, "name")

	regencies, err := h.repo.GetRegencies(name)
	if err != nil {
		return err
	}

	return interfaces.JSON(w, http.StatusOK, Response{
		Total: len(regencies),
		Data:  regencies,
	})
}

func (h *handler) HandlerDistricts(w http.ResponseWriter, r *http.Request) error {
	name := interfaces.GetQuery(r, "name")

	district, err := h.repo.GetDistricts(name)
	if err != nil {
		return err
	}

	return interfaces.JSON(w, http.StatusOK, Response{
		Total: len(district),
		Data:  district,
	})
}

func (h *handler) HandlerVillages(w http.ResponseWriter, r *http.Request) error {
	name := interfaces.GetQuery(r, "name")

	village, err := h.repo.GetVillages(name)
	if err != nil {
		return err
	}

	return interfaces.JSON(w, http.StatusOK, Response{
		Total: len(village),
		Data:  village,
	})
}
