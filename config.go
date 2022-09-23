package main

import (
	databaseinterface "github.com/fajarardiyanto/flt-go-database/interfaces"
	databaseLib "github.com/fajarardiyanto/flt-go-database/lib"
	"github.com/fajarardiyanto/flt-go-logger/interfaces"
	"github.com/fajarardiyanto/flt-go-logger/lib"
)

var (
	logger   interfaces.Logger
	database databaseinterface.SQL
)

func init() {
	logger = lib.NewLib()
	logger.Init("API Data Wilayah")

	db := databaseLib.NewLib()
	db.Init(logger)

	database = db.LoadSQLDatabase(databaseinterface.SQLConfig{
		Enable:        true,
		Driver:        "postgresql",
		Host:          "127.0.0.1",
		Port:          5432,
		Username:      "postgres",
		Password:      "postgres",
		Database:      "privy",
		AutoReconnect: true,
		StartInterval: 2,
	})
	if err := database.LoadSQL(); err != nil {
		logger.Error(err).Quit()
	}
}

func GetLogger() interfaces.Logger {
	return logger
}

func GetDBConn() databaseinterface.SQL {
	return database
}
