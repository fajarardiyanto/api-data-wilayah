package main

type ResultVillages struct {
	ID         int    `json:"id"`
	DistrictID int    `json:"district_id"`
	Name       string `json:"name"`
}

type ResultProvinces struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
