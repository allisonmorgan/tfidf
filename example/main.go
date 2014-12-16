package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/allisonmorgan/tfidf"
)

func ReadCSV(filepath string) ([][]string, error) {
	csvfile, err := os.Open(filepath)

	if err != nil {
		fmt.Printf("Unable to read csv: %v", err)
		return nil, err
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	fields, err := reader.ReadAll()

	return fields, err
}

func main() {
	frequency := tfidf.NewTermFrequencyStruct()

	subjects, err := ReadCSV("emailsubjects.csv")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, subject := range subjects {
			frequency.AddDocument(subject[0])
		}
		frequency.InverseDocumentFrequency()
		fmt.Println(frequency.InverseDocMap)
	}
}
