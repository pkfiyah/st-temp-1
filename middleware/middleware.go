package middleware

import (
	"fmt"
	"net/http"
	"pkfiyah/st-temp-1/utils"
)

// Used to ensure proper CSV, and respects constraints of API
// Eg. is square, has values
func InputValidation(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		records, err := utils.GetCSVContents(r)
		if err != nil {
			w.Write(fmt.Appendf([]byte{}, "error %s\n", err.Error()))
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
