package main

type ApiIndonesiaArea interface {
	GetProvinces(name string) ([]ResultProvinces, error)
	GetRegencies(string) ([]ResultVillages, error)
	GetCSV(string) ([][]string, error)
}
