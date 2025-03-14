package utils

import (
	"encoding/csv"
	"net/http"
)

// Input is expected to be a CSV File provided via the "file" form field in the request
func GetCSVContents(r *http.Request) ([][]string, error) {
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {

		return nil, err
	}
	return records, nil
}
