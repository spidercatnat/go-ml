package functions

// CSVRecord - In Go, we can se structs to enforce strict checks for each of the CSV fields
type CSVRecord struct {
	SepalLength float64
	SepalWidth  float64
	PetalLength float64
	PetalWidth  float64
	Species     string
	ParseError  error
}
