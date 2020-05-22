package functions

import (
	"encoding/csv"
	"os"
)

// ReadAllAtOnce ... Read all CSV records at once
func ReadAllAtOnce(f *os.File) [][]string {
	reader := csv.NewReader(f)
	// Setting to -1 allows variable fields. Will throw an error if data violates fields specified.
	reader.FieldsPerRecord = 5
	rawCSVData, err := reader.ReadAll()
	Check(err)
	return rawCSVData
}
