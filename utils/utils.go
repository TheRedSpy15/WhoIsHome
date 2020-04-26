package utils

import (
	"encoding/csv"
	"io"
	"os"
)

// ReadCsvFile - Returns an array from csv file. [0]=address, [1]=databaseName, [2]=username, [3]=password
func ReadCsvFile(filePath string) [4]string {
	// Load a csv file.
	f, _ := os.Open(filePath)
	var data [4]string

	// Create a new reader.
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		// add to array
		for value := range record {
			data[value] = record[value]
		}
	}

	return data
}
