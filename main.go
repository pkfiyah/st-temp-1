package main

import (
	"net/http"
	"pkfiyah/st-temp-1/handlers"
	"pkfiyah/st-temp-1/middleware"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {
	http.Handle("/echo", middleware.InputValidation(http.HandlerFunc(handlers.Echo)))
	http.Handle("/sum", middleware.InputValidation(http.HandlerFunc(handlers.Sum)))
	http.ListenAndServe(":8080", nil)
}
