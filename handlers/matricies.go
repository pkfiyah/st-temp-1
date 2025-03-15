package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"pkfiyah/st-temp-1/utils"
	"strconv"
	"strings"
)

// Provided Echo function code
func Echo(w http.ResponseWriter, r *http.Request) {
	records, err := utils.GetCSVContents(r)
	if err != nil {
		w.Write(fmt.Appendf([]byte{}, "error %s\n", err.Error()))
		return
	}

	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	fmt.Fprint(w, response)
}

func Invert(w http.ResponseWriter, r *http.Request) {
	records, err := utils.GetCSVContents(r)
	if err != nil {
		w.Write(fmt.Appendf([]byte{}, "error %s", err.Error()))
		return
	}

	var response string
	for i := range records {
		for j := i; j < len(records); j++ {
			if i != j {
				temp := records[i][j]
				records[i][j] = records[j][i]
				records[j][i] = temp
			}
		}
		response = fmt.Sprintf("%s%s\n", response, strings.Join(records[i], ","))
	}

	fmt.Fprint(w, response)
}

func Flatten(w http.ResponseWriter, r *http.Request) {
	records, err := utils.GetCSVContents(r)
	if err != nil {
		w.Write(fmt.Appendf([]byte{}, "error %s", err.Error()))
		return
	}

	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s,", response, strings.Join(row, ","))
	}

	fmt.Fprintf(w, "%s\n", response[:len(response)-1])
}

func Sum(w http.ResponseWriter, r *http.Request) {
	records, err := utils.GetCSVContents(r)
	if err != nil {
		w.Write(fmt.Appendf([]byte{}, "error %s\n", err.Error()))
		return
	}

	sum := 0
	for i, row := range records {
		for j := range records[i] {
			val, err := strconv.ParseInt(row[j], 10, 64)
			if err != nil {
				if errors.Is(err, strconv.ErrSyntax) {
					continue
				}
				w.Write(fmt.Appendf([]byte{}, "error %s\n", err.Error()))
				return
			}

			sum += int(val)
		}
	}

	fmt.Fprintf(w, "%d\n", sum)
}

func Multiply(w http.ResponseWriter, r *http.Request) {
	records, err := utils.GetCSVContents(r)
	if err != nil {
		w.Write(fmt.Appendf([]byte{}, "error %s\n", err.Error()))
		return
	}

	mult := 1
	for i, row := range records {
		for j := range records[i] {
			val, err := strconv.ParseInt(row[j], 10, 64)
			if err != nil {
				if errors.Is(err, strconv.ErrSyntax) {
					continue
				}
				w.Write(fmt.Appendf([]byte{}, "error %s\n", err.Error()))
				return
			}

			if val == 0 {
				mult = 0
				break
			}

			mult *= int(val)
		}
	}

	fmt.Fprintf(w, "%d\n", mult)
}
