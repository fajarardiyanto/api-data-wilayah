package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/regencies", SetMiddlewareJSON(GetRegencies))
	r.HandleFunc("/api/provinces", SetMiddlewareJSON(GetProvince))
	r.HandleFunc("/api/districts", SetMiddlewareJSON(GetDistricts))
	r.HandleFunc("/api/villages", SetMiddlewareJSON(GetVillages))

	http.ListenAndServe(":3000", r)
}

// Get Province
func GetProvince(w http.ResponseWriter, r *http.Request) {
	dataCSV, err := GetCSV("data/provinces.csv")
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err.Error())
	}

	for _, record := range dataCSV {
		data := map[string]interface{}{
			"id":   record[0],
			"name": record[1],
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(data)
	}
}

// Get Kab/Kota
func GetRegencies(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	dataCSV, err := GetCSV("data/regencies.csv")
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err.Error())
	}

	for _, record := range dataCSV {
		if id == record[1] {
			data := map[string]interface{}{
				"id":          record[0],
				"province_id": record[1],
				"name":        record[2],
			}

			w.WriteHeader(200)
			json.NewEncoder(w).Encode(data)
		}
	}
}

// Get Kecamatan
func GetDistricts(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	dataCSV, err := GetCSV("data/districts.csv")
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err.Error())
	}

	for _, record := range dataCSV {
		if id == record[1] {
			data := map[string]interface{}{
				"id":         record[0],
				"regency_id": record[1],
				"name":       record[2],
			}

			w.WriteHeader(200)
			json.NewEncoder(w).Encode(data)
		}
	}
}

// Get Kelurahan
func GetVillages(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	dataCSV, err := GetCSV("data/villages.csv")
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err.Error())
	}

	for _, record := range dataCSV {
		if id == record[1] {
			data := map[string]interface{}{
				"id":          record[0],
				"district_id": record[1],
				"name":        record[2],
			}

			w.WriteHeader(200)
			json.NewEncoder(w).Encode(data)
		}
	}
}

func GetCSV(file string) ([][]string, error) {
	csvfile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	rd := csv.NewReader(csvfile)

	record, err := rd.ReadAll()
	if err == io.EOF {
		return nil, err
	}
	if err != nil {
		log.Fatal(err)
	}

	return record, nil
}
