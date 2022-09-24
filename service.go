package main

import (
	"encoding/csv"
	"github.com/fajarardiyanto/flt-go-utils/parser"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"
)

type service struct {
	sync.Mutex
}

func NewService() ApiIndonesiaArea {
	return &service{}
}

// GetProvinces service get provinces
func (s *service) GetProvinces(name string) ([]ResultProvinces, error) {
	dataCSV, err := s.GetCSV("data/provinces.csv")
	if err != nil {
		return nil, err
	}

	var datas []ResultProvinces
	name = strings.ToUpper(name)
	re, _ := regexp.Compile(name)

	var wg sync.WaitGroup

	for _, record := range dataCSV {

		wg.Add(1)
		go func(wg *sync.WaitGroup, record []string) {
			defer wg.Done()

			result := re.FindAllString(record[2], -1)
			if len(result) != 0 {
				s.Lock()
				postal := strings.Split(record[5], ",")
				s.Unlock()

				s.Lock()
				data := ResultProvinces{
					ID:   parser.StrToInt(record[0]),
					Name: record[2],
					MetaData: MetaData{
						Latitude:  record[3],
						Longitude: record[4],
						Postal:    postal,
					},
				}
				s.Unlock()

				s.Lock()
				datas = append(datas, data)
				s.Unlock()
			}
		}(&wg, record)
		wg.Wait()

	}

	return datas, nil
}

// GetRegencies service get regencies / cities
func (s *service) GetRegencies(name string) ([]ResultRegencies, error) {
	dataCSV, err := s.GetCSV("data/regencies.csv")
	if err != nil {
		return nil, err
	}

	var datas []ResultRegencies
	name = strings.ToUpper(name)
	re, _ := regexp.Compile(name)

	var wg sync.WaitGroup

	for _, record := range dataCSV {

		wg.Add(1)
		go func(wg *sync.WaitGroup, record []string) {
			defer wg.Done()

			result := re.FindAllString(record[2], -1)
			if len(result) != 0 {
				s.Lock()
				postal := strings.Split(record[5], ",")
				s.Unlock()

				s.Lock()
				data := ResultRegencies{
					ID:     parser.StrToInt(record[0]),
					ProvID: parser.StrToInt(record[1]),
					Name:   record[2],
					MetaData: MetaData{
						Latitude:  record[3],
						Longitude: record[4],
						Postal:    postal,
					},
				}
				s.Unlock()

				s.Lock()
				datas = append(datas, data)
				s.Unlock()
			}
		}(&wg, record)
		wg.Wait()

	}

	return datas, nil
}

// GetDistricts service get regencies / cities
func (s *service) GetDistricts(name string) ([]ResultDistricts, error) {
	dataCSV, err := s.GetCSV("data/districts.csv")
	if err != nil {
		return nil, err
	}

	var datas []ResultDistricts
	name = strings.ToUpper(name)
	re, _ := regexp.Compile(name)

	var wg sync.WaitGroup

	for _, record := range dataCSV {

		wg.Add(1)
		go func(wg *sync.WaitGroup, record []string) {
			defer wg.Done()

			result := re.FindAllString(record[2], -1)
			if len(result) != 0 {
				s.Lock()
				postal := strings.Split(record[5], ",")
				s.Unlock()

				s.Lock()
				data := ResultDistricts{
					ID:          parser.StrToInt(record[0]),
					RegenciesID: parser.StrToInt(record[1]),
					Name:        record[2],
					MetaData: MetaData{
						Latitude:  record[3],
						Longitude: record[4],
						Postal:    postal,
					},
				}
				s.Unlock()

				s.Lock()
				datas = append(datas, data)
				s.Unlock()
			}
		}(&wg, record)
		wg.Wait()

	}

	return datas, nil
}

// GetVillages service get regencies / cities
func (s *service) GetVillages(name string) ([]ResultVillages, error) {
	dataCSV, err := s.GetCSV("data/villages.csv")
	if err != nil {
		return nil, err
	}

	var datas []ResultVillages
	name = strings.ToUpper(name)
	re, _ := regexp.Compile(name)

	var wg sync.WaitGroup

	for _, record := range dataCSV {

		wg.Add(1)
		go func(wg *sync.WaitGroup, record []string) {
			defer wg.Done()

			result := re.FindAllString(record[2], -1)
			if len(result) != 0 {
				s.Lock()
				postal := strings.Split(record[5], ",")
				s.Unlock()

				s.Lock()
				data := ResultVillages{
					ID:         parser.StrToInt(record[0]),
					DistrictId: parser.StrToInt(record[1]),
					Name:       record[2],
					MetaData: MetaData{
						Postal: postal,
					},
				}
				s.Unlock()

				s.Lock()
				datas = append(datas, data)
				s.Unlock()
			}
		}(&wg, record)
		wg.Wait()

	}

	return datas, nil
}

func (s *service) GetCSV(file string) ([][]string, error) {
	csvfile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	rd := csv.NewReader(csvfile)

	var wg sync.WaitGroup
	var errs []error
	resChan := make(chan [][]string)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		record, err := rd.ReadAll()
		if err == io.EOF {
			s.Lock()
			errs = append(errs, err)
			s.Unlock()
		}
		if err != nil {
			s.Lock()
			errs = append(errs, err)
			s.Unlock()
		}

		s.Lock()
		resChan <- record
		s.Unlock()
	}(&wg)

	result := <-resChan
	wg.Wait()

	if len(errs) != 0 {
		err = errs[0]
	}

	return result, err
}
