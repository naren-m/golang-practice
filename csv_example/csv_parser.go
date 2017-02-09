package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Source Documentation https://golang.org/pkg/encoding/csv/
// Playground URL https://play.golang.org/p/VHj-F2GOuU
// TODO: Handle errors
func main() {
	file, err := os.Open("csv_example/data/sample.csv")
	if err != nil {
		fmt.Print(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	if err != nil {
		fmt.Print(err)
	}

	for i, line := range lines {
		if i == 0 {
			continue
		}
		fmt.Printf("Name: %s %s, SSN: %s \n", line[0], line[1], line[2])
	}

}
