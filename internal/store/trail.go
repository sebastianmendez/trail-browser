package store

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/sebastianmendez/trail-browser/internal/store/model"
)

var (
	CsvPath = "assets/BoulderTrailHeads.csv"
)

// this layer will take responsability of handling raw data and convert it to a model that the app knows
func List(ctx context.Context, params map[string]string) ([]model.Trail, error) {
	applyNameFilter := false
	applyParkFilter := false
	applyBothFilters := false
	parkSpacesFilter := -1
	if params[model.FaccessName] != "" && params[model.FparkSpaces] != "" {
		applyBothFilters = true
		var err error
		parkSpacesFilter, err = strconv.Atoi(params[model.FparkSpaces])
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
	}
	if params[model.FaccessName] != "" {
		applyNameFilter = true
	}
	if params[model.FparkSpaces] != "" {
		applyParkFilter = true
		var err error
		parkSpacesFilter, err = strconv.Atoi(params[model.FparkSpaces])
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
	}
	file, err := os.Open(CsvPath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	defer file.Close()
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	res := []model.Trail{}
	for i := range records {
		if i > 0 { // skipping headers
			trail, err := model.ParseEntry(records[i])
			if err != nil {
				log.Default().Printf("Error parsing entry, skipping: %v", err)
			}
			if applyBothFilters {
				if strings.Contains(strings.ToLower(trail.AccessName), strings.ToLower(params[model.FaccessName])) && trail.ParkSpaces == int32(parkSpacesFilter) && parkSpacesFilter > -1 {
					res = append(res, trail)
				}
				continue
			}
			if applyNameFilter {
				if strings.Contains(strings.ToLower(trail.AccessName), strings.ToLower(params[model.FaccessName])) {
					res = append(res, trail)
				}
				continue
			}
			if applyParkFilter {
				if trail.ParkSpaces == int32(parkSpacesFilter) && parkSpacesFilter > -1 {
					res = append(res, trail)
				}
				continue
			}
			res = append(res, trail)
		}
	}
	return res, nil
}
