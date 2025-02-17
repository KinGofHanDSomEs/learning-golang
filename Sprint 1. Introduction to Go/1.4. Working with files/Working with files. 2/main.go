package main

import (
	"bufio"
	"os"
)

func LineByNum(inputFilename string, lineNum int) string {
	file, err := os.Open(inputFilename)
	if err != nil {
		return ""
	}
	fileScanner := bufio.NewScanner(file)
	k := 0
	for fileScanner.Scan() {
		if k == lineNum {
			return fileScanner.Text()
		}
		k++
	}
	return ""
}
