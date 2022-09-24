package main

import (
	"github.com/fajarardiyanto/flt-go-router/interfaces"
	"math"
	"net/http"
)

type handler struct {
	repo ApiIndonesiaArea
}

func NewHandler(repo ApiIndonesiaArea) *handler {
	return &handler{repo}
}

func (h *handler) HandlerInfo(w http.ResponseWriter, r *http.Request) error {
	host := r.Host
	data := []Info{
		{
			Name: "provinces",
			Link: host + "/provinces",
		},
		{
			Name: "regencies",
			Link: host + "/regencies",
		},
		{
			Name: "districts",
			Link: host + "/districts",
		},
		{
			Name: "villages",
			Link: host + "/villages",
		},
	}
	return interfaces.JSON(w, http.StatusOK, data)
}

func (h *handler) HandlerProvinces(w http.ResponseWriter, r *http.Request) error {
	page := GetQueryPage(r)
	offset := GetQueryOffset(r)

	prov, total, err := h.repo.GetProvinces(GetQueryName(r), page, offset)
	if err != nil {
		return err
	}

	var totalPage int
	if page != 0 {
		totalPage = int(math.Ceil(float64(total) / float64(offset)))
	}

	return interfaces.JSON(w, http.StatusOK, Response{
		Total:     total,
		TotalPage: totalPage,
		Data:      prov,
	})
}

func (h *handler) HandlerRegencies(w http.ResponseWriter, r *http.Request) error {
	page := GetQueryPage(r)
	offset := GetQueryOffset(r)

	regencies, total, err := h.repo.GetRegencies(GetQueryName(r), page, offset)
	if err != nil {
		return err
	}

	var totalPage int
	if page != 0 {
		totalPage = int(math.Ceil(float64(total) / float64(offset)))
	}

	return interfaces.JSON(w, http.StatusOK, Response{
		Total:     total,
		TotalPage: totalPage,
		Data:      regencies,
	})
}

func (h *handler) HandlerDistricts(w http.ResponseWriter, r *http.Request) error {
	page := GetQueryPage(r)
	offset := GetQueryOffset(r)

	district, total, err := h.repo.GetDistricts(GetQueryName(r), page, offset)
	if err != nil {
		return err
	}

	var totalPage int
	if page != 0 {
		totalPage = int(math.Ceil(float64(total) / float64(offset)))
	}

	return interfaces.JSON(w, http.StatusOK, Response{
		Total:     total,
		TotalPage: totalPage,
		Data:      district,
	})
}

func (h *handler) HandlerVillages(w http.ResponseWriter, r *http.Request) error {
	page := GetQueryPage(r)
	offset := GetQueryOffset(r)

	village, total, err := h.repo.GetVillages(GetQueryName(r), page, offset)
	if err != nil {
		return err
	}

	var totalPage int
	if page != 0 {
		totalPage = int(math.Ceil(float64(total) / float64(offset)))
	}

	return interfaces.JSON(w, http.StatusOK, Response{
		Total:     total,
		TotalPage: totalPage,
		Data:      village,
	})
}
