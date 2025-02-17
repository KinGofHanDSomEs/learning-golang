package main

import (
	"os"
)

func WriteToLogFile(message string, fileName string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	if _, err := file.Write([]byte(message)); err != nil {
		return err
	}
	return nil
}
