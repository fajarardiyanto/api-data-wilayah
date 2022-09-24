package main

type ApiIndonesiaArea interface {
	GetProvinces(string, int, int) ([]ResultProvinces, int, error)
	GetRegencies(string, int, int) ([]ResultRegencies, int, error)
	GetDistricts(string, int, int) ([]ResultDistricts, int, error)
	GetVillages(string, int, int) ([]ResultVillages, int, error)
	GetCSV(string) ([][]string, error)
}
