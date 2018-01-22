package main

import (
	"bytes"
	"encoding/csv"
	"log"
	"os"
)

func main() {

	// Open the diabetes dataset file.
	f, err := os.Open("../../data/diabetes.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(bytes.NewReader(b.Bytes()))
	reader.FieldsPerRecord = 11

	// Read in all of the CSV records
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Sequentially move the rows into a slice of floats.

	// Loop over the columns.

	// Add the float values to a slice of floats.

	// Form the matrix.

	// Get the first 10 rows.

	// As a sanity check, output the rows to standard out.
}
