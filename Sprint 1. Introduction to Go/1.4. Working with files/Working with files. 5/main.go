package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
	var result []string
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		return result, err
	}
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		logDate, err := time.Parse("02.01.2006", strings.Split(line, " ")[0])
		if err != nil {
			continue
		}
		if logDate.After(start.AddDate(0, 0, -1)) && logDate.Before(end.AddDate(0, 0, 1)) {
			result = append(result, line)
		}
	}
	if len(result) == 0 {
		return result, fmt.Errorf("file without log-files")
	}
	return result, nil
}