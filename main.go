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
	http.Handle("/invert", middleware.InputValidation(http.HandlerFunc(handlers.Invert)))
	http.Handle("/flatten", middleware.InputValidation(http.HandlerFunc(handlers.Flatten)))
	http.Handle("/sum", middleware.InputValidation(http.HandlerFunc(handlers.Sum)))
	http.Handle("/multiply", middleware.InputValidation(http.HandlerFunc(handlers.Multiply)))
	http.ListenAndServe(":8080", nil)
}
