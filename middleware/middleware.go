package middleware

import (
	"encoding/csv"
	"fmt"
	"net/http"
)

func InputValidation(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		defer file.Close()
		records, err := csv.NewReader(file).ReadAll()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		if len(records) == 0 {
			w.Write([]byte("no data found in provided csv\n"))
			return
		}
		if len(records) != len(records[0]) {
			w.Write([]byte("only square matrices are accepted\n"))
			return
		}
		h.ServeHTTP(w, r)
	})
}
