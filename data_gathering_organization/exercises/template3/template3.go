package main

import (
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {

	// Pull in the CSV file.
	irisFile, err := os.Open("../../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	irisDF := dataframe.ReadCSV(irisFile)

	// Filter the dataset into three separate dataframes,
	// each corresponding to one of the Iris species.

	// Save each of the species dataframe to a file.
}
