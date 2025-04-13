package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ReadCSV(file string) (<-chan []string, error) {
	fl, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	out := make(chan []string)
	go func() {
		defer close(out)
		defer fl.Close()
		reader := csv.NewReader(fl)
		for {
			row, err := reader.Read()
			if err != nil {
				return
			}
			out <- row
		}
	}()

	return out, nil
}

func main() {
	file := "test.csv"
	var rows [][]string
	out, err := ReadCSV(file)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		select {
		case row, ok := <-out:
			if !ok {
				fmt.Println(rows)
				return

			}
			rows = append(rows, row)
		}
	}
}
