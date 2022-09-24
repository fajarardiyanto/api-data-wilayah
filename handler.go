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

	prov, total, totalPage, err := h.repo.GetProvinces(GetQueryName(r), page, offset)
	if err != nil {
		return err
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

	regencies, total, totalPage, err := h.repo.GetRegencies(GetQueryName(r), page, offset)
	if err != nil {
		return err
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

	district, total, totalPage, err := h.repo.GetDistricts(GetQueryName(r), page, offset)
	if err != nil {
		return err
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

	village, total, totalPage, err := h.repo.GetVillages(GetQueryName(r), page, offset)
	if err != nil {
		return err
	}

	return interfaces.JSON(w, http.StatusOK, Response{
		Total:     total,
		TotalPage: totalPage,
		Data:      village,
	})
}
