package main

import (
	"encoding/csv"
	"fmt"
	"strconv"

	"github.com/spidercatnat/go-ml/1-handling-data/functions"
)

func main() {
	f := functions.OpenFile("./data/test.csv")
	defer f.Close()

	// Read in the CSV records
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	functions.Check(err)

	// Get the maximum value in the integer column
	var intMax int
	for _, record := range records {
		// Parse the integer value
		intVal, err := strconv.Atoi(record[0])
		// In contrast to Python, Go will throw an error if our CSV data is does not fit the specified format.
		functions.Check(err)
		if intVal > intMax {
			intMax = intVal
		}
	}

	fmt.Println(intMax)
}
