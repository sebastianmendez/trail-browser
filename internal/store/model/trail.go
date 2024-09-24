package model

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

// this layer will have the object representation of every element of the domain and its associated functions to make it available
// for the rest of the app that will use it
const (
	FparkSpaces = "park_spaces"
	FaccessName = "access_name"
)

type Trail struct {
	Fid        int32  `json:"fid"`
	Restrooms  bool   `json:"restrooms"`
	Picninc    bool   `json:"picninc"`
	Fishing    bool   `json:"fishing"`
	Aka        string `json:"aka"`
	AccessType string `json:"access_type"`
	AccessID   string `json:"access_id"`
	Class      string `json:"class"`
	Address    string `json:"address"`
	Fee        bool   `json:"fee"`
	BikeRack   bool   `json:"bike_rack"`
	BikeTrail  bool   `json:"bike_trail"`
	DogTube    int32  `json:"dog_tube"`
	Grills     bool   `json:"grills"`
	TrashCans  int32  `json:"trash_cans"`
	ParkSpaces int32  `json:"park_spaces"`
	ADAsurface string `json:"ADA_surface"`
	ADAtoilet  bool   `json:"ADA_toilet"`
	ADAfishing bool   `json:"ADA_fishing"`
	ADAcamping bool   `json:"ADA_camping"`
	ADApicnic  bool   `json:"ADA_picnic"`
	ADAtrail   string `json:"ADA_trail"`
	ADAparking string `json:"ADA_parking"`
	ADAfacilit bool   `json:"ADA_facilit"`
	ADAfacName string `json:"ADA_fac_name"`
	HorseTrail string `json:"horse_trail"`
	DateFrom   string `json:"date_from"`
	DateTo     string `json:"date_to"`
	RecycleBin bool   `json:"recycle_bin"`
	DogCompost bool   `json:"dog_compost"`
	AccessName string `json:"access_name"`
	THLeash    bool   `json:"TH_leash"`
}

func ParseEntry(record []string) (Trail, error) {
	if len(record) != 32 {
		return Trail{}, errors.New("invalid structure")
	}
	fid, err := strconv.Atoi(record[0])
	if err != nil {
		log.Default().Printf("fid parse failed: %v", err)
	}
	dogTube, err := strconv.Atoi(record[12])
	if err != nil {
		log.Default().Printf("dog tube parse failed: %v", err)
	}
	trashCans, err := strconv.Atoi(record[14])
	if err != nil {
		log.Default().Printf("trashCans parse failed: %v", err)
	}
	parkSpaces, err := strconv.Atoi(record[15])
	if err != nil {
		log.Default().Printf("parkSpaces parse failed: %v", err)
	}
	trail := Trail{}
	trail.Fid = int32(fid)
	trail.Restrooms = parseBoolean(record[1])
	trail.Picninc = parseBoolean(record[2])
	trail.Fishing = parseBoolean(record[3])
	trail.Aka = record[4]
	trail.AccessType = record[5]
	trail.AccessID = record[6]
	trail.Class = record[7]
	trail.Address = record[8]
	trail.Fee = parseBoolean(record[9])
	trail.BikeRack = parseBoolean(record[10])
	trail.BikeTrail = parseBoolean(record[11])
	trail.DogTube = int32(dogTube)
	trail.Grills = parseBoolean(record[13])
	trail.TrashCans = int32(trashCans)
	trail.ParkSpaces = int32(parkSpaces)
	trail.ADAsurface = record[16]
	trail.ADAtoilet = parseBoolean(record[17])
	trail.ADAfishing = parseBoolean(record[18])
	trail.ADAcamping = parseBoolean(record[19])
	trail.ADApicnic = parseBoolean(record[20])
	trail.ADAtrail = record[21]
	trail.ADAparking = record[22]
	trail.ADAfacilit = parseBoolean(record[23])
	trail.ADAfacName = record[24]
	trail.HorseTrail = record[25]
	trail.DateFrom = record[26]
	trail.DateTo = record[27]
	trail.RecycleBin = parseBoolean(record[28])
	trail.DogCompost = parseBoolean(record[29])
	trail.AccessName = record[30]
	trail.THLeash = parseBoolean(record[31])
	return trail, nil
}

func parseBoolean(value string) bool {
	return strings.ToLower(value) == "yes"
}
