package main

import (
	"github.com/fajarardiyanto/flt-go-router/interfaces"
	"github.com/fajarardiyanto/flt-go-utils/parser"
	"net/http"
)

func GetQueryName(r *http.Request) string {
	return interfaces.GetQuery(r, "name")
}

func GetQueryPage(r *http.Request) int {
	return parser.StrToInt(interfaces.GetQuery(r, "page"))
}

func GetQueryOffset(r *http.Request) int {
	return parser.StrToInt(interfaces.GetQuery(r, "offset"))
}
