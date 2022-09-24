package main

type ApiIndonesiaArea interface {
	GetProvinces(name string) ([]ResultProvinces, error)
	GetRegencies(string) ([]ResultRegencies, error)
	GetDistricts(name string) ([]ResultDistricts, error)
	GetVillages(name string) ([]ResultVillages, error)
	GetCSV(string) ([][]string, error)
}
