package main

import (
	"github.com/fajarardiyanto/flt-go-logger/interfaces"
	"github.com/fajarardiyanto/flt-go-logger/lib"
)

var logger interfaces.Logger

func init() {
	logger = lib.NewLib()
	logger.Init("API Data Wilayah")
}

func GetLogger() interfaces.Logger {
	return logger
}
