package main

import (
	"fmt"
	"log"

	"github.com/go-gota/gota/dataframe"
	"github.com/spidercatnat/go-ml/1-handling-data/functions"
)

func main() {
	irisFile := functions.OpenFile("./data/iris.csv")
	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	irisDF := dataframe.ReadCSV(irisFile)

	// As a sanity check, display the records to stdout.
	// Gota will format the dataframe for pretty printing.
	// fmt.Println(irisDF)

	// Create a filter for the dataframe
	filter := dataframe.F{
		Colname:    "species",
		Comparator: "==",
		Comparando: "versicolor",
	}

	// Filter the dataframe where species is "Iris-versicolor"
	versicolorDF := irisDF.Filter(filter)
	if versicolorDF.Err != nil {
		log.Fatal(versicolorDF.Err)
	}
	// fmt.Println(versicolorDF)

	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"})
	// fmt.Println(versicolorDF)

	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"}).Subset([]int{0, 1, 2})

	fmt.Println(versicolorDF)
}
