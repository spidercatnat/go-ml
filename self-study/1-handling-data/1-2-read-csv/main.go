package main

import (
	"fmt"

	"github.com/spidercatnat/go-ml/1-handling-data/functions"
)

// Using the functions defined below, load CSV and read out all records
func main() {
	// Load CSV
	f := functions.OpenFile("./data/iris.csv")
	defer f.Close()

	// Read all records iteratively
	fmt.Println(functions.ReadAllIterative(f))

	// Read all records at once
	// fmt.Println(functions.ReadAllAtOnce(f))
}
