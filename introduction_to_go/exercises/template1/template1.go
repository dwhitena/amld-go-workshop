package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

func main() {

	// Open the CSV.
	f, err := os.Open("../../data/example_messy.csv")
	if err != nil {
		err = errors.Wrap(err, "Could not open CSV")
		log.Fatal(err)
	}

	// Read in the CSV records
	r := csv.NewReader(bufio.NewReader(f))
	records, err := r.ReadAll()
	if err != nil {
		err = errors.Wrap(err, "Could not parse CSV")
		log.Fatal(err)
	}

	// Get the maximum value in the integer column.
	var intMax int
	for _, record := range records {
		intVal, err := strconv.Atoi(record[0])
		if err != nil {

			// Handle this error related to missing data.
		}
		if intVal > intMax {
			intMax = intVal
		}
	}

	// Print the maxium value
	fmt.Println(intMax)
}
