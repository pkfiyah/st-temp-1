package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Provided Echo function code
func Echo(w http.ResponseWriter, r *http.Request) {
	records, err := getCSVContents(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response)
}

func Invert(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func Flatten(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func Sum(w http.ResponseWriter, r *http.Request) {
	records, err := getCSVContents(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	var sum int64
	for i, row := range records {
		for j, _ := range records[i] {
			val, err := strconv.ParseInt(row[j], 10, 64)
			if err != nil {
				w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
				return
			}
			sum += val
		}
	}
	fmt.Fprintf(w, "%d\n", sum)
}

func Multiply(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// Input is expected to be a CSV File provided via the "file" form field in the request
func getCSVContents(r *http.Request) ([][]string, error) {
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
