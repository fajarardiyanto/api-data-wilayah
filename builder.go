package main

type ApiIndonesiaArea interface {
	GetProvinces(string, int, int) ([]ResultProvinces, int, int, error)
	GetRegencies(string, int, int, int) ([]ResultRegencies, int, int, error)
	GetDistricts(string, int, int, int) ([]ResultDistricts, int, int, error)
	GetVillages(string, int, int, int) ([]ResultVillages, int, int, error)
	GetCSV(string) ([][]string, error)
}
