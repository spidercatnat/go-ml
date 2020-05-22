package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// In contrast to Python, Go will throw an error if our CSV data is corrupted.

func main() {
	// Open the CSV
	f, err := os.Open("./load-csv/test.csv")
	check(err)

	// Read in the CSV records
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	check(err)

	// Get the maximum value in the integer column
	var intMax int
	for _, record := range records {
		// Parse the integer value
		intVal, err := strconv.Atoi(record[0])
		check(err)
		if intVal > intMax {
			intMax = intVal
		}
	}

	fmt.Println(intMax)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
