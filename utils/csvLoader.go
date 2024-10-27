package utils

import (
	"challenge2016/region"
	"encoding/csv"
	"os"
)

func LoadRegions(filePath string) ([]region.Region, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var regions []region.Region
	for _, record := range records {
		regions = append(regions, region.Region{
			Country: record[0],
			State:   record[1],
			City:    record[2],
		})
	}
	return regions, nil
}
