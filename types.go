package main

type Response struct {
	Total     int         `json:"total"`
	TotalPage int         `json:"total_page,omitempty"`
	Data      interface{} `json:"data"`
}

type Info struct {
	Name string `json:"name"`
	Link string `json:"link"`
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
	Latitude  string   `json:"latitude,omitempty"`
	Longitude string   `json:"longitude,omitempty"`
	Postal    []string `json:"postal"`
}
