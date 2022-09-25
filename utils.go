package main

import (
	"github.com/fajarardiyanto/flt-go-router/interfaces"
	"github.com/fajarardiyanto/flt-go-utils/parser"
	"net/http"
)

func GetQueryName(r *http.Request) string {
	return interfaces.GetQuery(r, "name")
}

func GetQueryProvID(r *http.Request) int {
	return parser.StrToInt(interfaces.GetQuery(r, "prov_id"))
}

func GetQueryRegenciesID(r *http.Request) int {
	return parser.StrToInt(interfaces.GetQuery(r, "regencies_id"))
}

func GetQueryDistrictID(r *http.Request) int {
	return parser.StrToInt(interfaces.GetQuery(r, "district_id"))
}

func GetQueryPage(r *http.Request) int {
	return parser.StrToInt(interfaces.GetQuery(r, "page"))
}

func GetQueryOffset(r *http.Request) int {
	return parser.StrToInt(interfaces.GetQuery(r, "offset"))
}
