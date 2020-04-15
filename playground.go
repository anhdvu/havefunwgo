package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type record struct {
	field0, field1, field2, field3 string
}

func main() {
	// Data
	record1 := record{"1", "2", "3", "4"}
	record2 := record{"a", "b", "c", "d"}

	records := []record{record1, record2}
	result := [][]string{}

	for _, record := range records {
		temp := make([]string, 4)
		assign(temp, record)
		result = append(result, temp)
	}

	// CLI flags

	outFlag := flag.String("o", "output", "Filename of the output file")
	flag.Parse()

	// Create and write to file
	ext := ".csv"
	f, err := os.Create(*outFlag + ext)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	w.WriteAll(result)

}

func assign(sa []string, rc record) {
	sa[0] = rc.field0
	sa[1] = rc.field1
	sa[2] = rc.field2
	sa[3] = rc.field3
}
