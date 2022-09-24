package main

type Response struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

type ResultProvinces struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	MetaData MetaData `json:"meta_data"`
}

type ResultRegencies struct {
	ID       int      `json:"id"`
	ProvID   int      `json:"prov_id"`
	Name     string   `json:"name"`
	MetaData MetaData `json:"meta_data"`
}

type ResultDistricts struct {
	ID          int      `json:"id"`
	RegenciesID int      `json:"regencies_id"`
	Name        string   `json:"name"`
	MetaData    MetaData `json:"meta_data"`
}

type ResultVillages struct {
	ID         int      `json:"id"`
	DistrictId int      `json:"district_id"`
	Name       string   `json:"name"`
	MetaData   MetaData `json:"meta_data"`
}

type MetaData struct {
	Latitude  string   `json:"latitude"`
	Longitude string   `json:"longitude"`
	Postal    []string `json:"postal"`
}
