package main

import (
	"encoding/csv"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"
)

type pooling struct {
	sync.Mutex
}

func NewPooling() *pooling {
	return &pooling{}
}

func (c *pooling) GetRegencies(name string) ([]ResultVillages, error) {
	dataCSV, err := c.GetCSV("data/regencies.csv")
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

				c.Lock()
				data := ResultVillages{
					ID:         record[0],
					DistrictID: record[1],
					Name:       record[2],
				}
				c.Unlock()

				c.Lock()
				datas = append(datas, data)
				c.Unlock()
			}
		}(&wg, record)
		wg.Wait()

	}

	return datas, nil
}

func (c *pooling) GetCSV(file string) ([][]string, error) {
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
			c.Lock()
			errs = append(errs, err)
			c.Unlock()
		}
		if err != nil {
			c.Lock()
			errs = append(errs, err)
			c.Unlock()
		}

		c.Lock()
		resChan <- record
		c.Unlock()
	}(&wg)

	result := <-resChan
	wg.Wait()

	if len(errs) != 0 {
		err = errs[0]
	}

	return result, err
}
