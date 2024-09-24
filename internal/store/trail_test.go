package store_test

import (
	"context"
	"os"
	"testing"

	"github.com/sebastianmendez/trail-browser/internal/store"
	"github.com/sebastianmendez/trail-browser/internal/store/model"
	"github.com/stretchr/testify/assert"
)

var mockCSVData = `
fid,restrooms,picninc,fishing,aka,access_type,access_id,class,address,fee,bike_rack,bike_trail,dog_tube,grills,trash_cans,park_spaces,ADA_surface,ADA_toilet,ADA_fishing,ADA_camping,ADA_picnic,ADA_trail,ADA_parking,ADA_facilit,ADA_fac_name,horse_trail,date_from,date_to,recycle_bin,dog_compost,access_name,TH_leash
1,yes,no,yes,Trail A,road,001,Class A,1234 Elm St,yes,yes,yes,3,yes,5,50,concrete,yes,yes,no,yes,gravel,yes,no,ADA Facility A,Horse Trail 1,2023-01-01,2023-12-31,yes,no,North Entrance,yes
2,no,yes,no,Trail B,footpath,002,Class B,5678 Oak St,no,no,yes,0,no,10,20,dirt,no,no,no,no,gravel,no,yes,ADA Facility B,Horse Trail 2,2023-05-01,2023-11-30,no,yes,South Entrance,no
3,yes,yes,yes,Trail C,road,003,Class C,9876 Pine St,yes,no,no,5,no,7,100,concrete,yes,no,yes,yes,gravel,no,no,ADA Facility C,Horse Trail 3,2023-03-01,2023-10-31,yes,yes,East Entrance,yes
`

func createTempCSVFile(t *testing.T, content string) string {
	tempFile, err := os.CreateTemp("", "trailheads-*.csv")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	if _, err := tempFile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write mock CSV data: %v", err)
	}

	if err := tempFile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	return tempFile.Name()
}

func TestList_NoFilters(t *testing.T) {
	csvPath := createTempCSVFile(t, mockCSVData)
	defer os.Remove(csvPath) // Cleanup after test

	store.CsvPath = csvPath

	params := map[string]string{}

	ctx := context.Background()
	trails, err := store.List(ctx, params)

	assert.NoError(t, err)
	assert.Len(t, trails, 3) // Should return both trails since no filter is applied
}

func TestList_AccessNameFilter(t *testing.T) {
	csvPath := createTempCSVFile(t, mockCSVData)
	defer os.Remove(csvPath)
	store.CsvPath = csvPath
	params := map[string]string{
		model.FaccessName: "North Entrance",
	}
	ctx := context.Background()
	trails, err := store.List(ctx, params)
	assert.NoError(t, err)
	assert.Len(t, trails, 1)
	assert.Equal(t, "North Entrance", trails[0].AccessName)
}

func TestList_ParkSpacesFilter(t *testing.T) {
	csvPath := createTempCSVFile(t, mockCSVData)
	defer os.Remove(csvPath)
	store.CsvPath = csvPath
	params := map[string]string{
		model.FparkSpaces: "50",
	}
	ctx := context.Background()
	trails, err := store.List(ctx, params)
	assert.NoError(t, err)
	assert.Len(t, trails, 1)
	assert.Equal(t, int32(50), trails[0].ParkSpaces)
}

func TestList_AccessNameAndParkSpacesFilter(t *testing.T) {
	csvPath := createTempCSVFile(t, mockCSVData)
	defer os.Remove(csvPath)
	store.CsvPath = csvPath
	params := map[string]string{
		model.FaccessName: "North Entrance",
		model.FparkSpaces: "50",
	}
	ctx := context.Background()
	trails, err := store.List(ctx, params)
	assert.NoError(t, err)
	assert.Len(t, trails, 1)
	assert.Equal(t, "North Entrance", trails[0].AccessName)
	assert.Equal(t, int32(50), trails[0].ParkSpaces)
}
