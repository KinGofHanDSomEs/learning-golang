package main

import (
	"io"
	"os"
)

func CopyFilePart(inputFilename, outFileName string, startPos int) error {
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	inputFile.Seek(int64(startPos), 0)

	outFile, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, inputFile)
	if err != nil {
		return err
	}
	return nil
}
