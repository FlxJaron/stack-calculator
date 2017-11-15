package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialise New Calculator
	calc := NewCalculator()

	// Start the web server on port 8080
	log.Fatal(http.ListenAndServe(":8080", calc.Router))
}
