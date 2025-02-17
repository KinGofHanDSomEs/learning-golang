package main

import (
	"fmt"
	"os"
	"strconv"
)

func run() error {
	input := os.Args
	if len(input) != 4 {
		return fmt.Errorf("invalid values")
	}
	width, err := strconv.Atoi(input[1])
	if err != nil {
		return fmt.Errorf("width is not integer")
	}
	height, err := strconv.Atoi(input[2])
	if err != nil {
		return fmt.Errorf("height is not integer")
	}
	percent, err := strconv.Atoi(input[3])
	if err != nil {
		return fmt.Errorf("percent is not integer")
	}
	result := fmt.Sprintf("%dx%d %d%%", width, height, percent)
	err = os.WriteFile("config.txt", []byte(result), 0644)
	if err != nil {
		return fmt.Errorf("file recording error")
	}
	return nil
}
