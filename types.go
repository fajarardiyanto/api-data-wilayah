package main

type ResultVillages struct {
	ID         string `json:"id" gorm:"column:id"`
	DistrictID string `json:"district_id" gorm:"column:district_id"`
	Name       string `json:"name" gorm:"column:name"`
}
