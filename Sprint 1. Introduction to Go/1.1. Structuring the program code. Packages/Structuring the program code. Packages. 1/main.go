package main

import (
	"fmt"
)

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func NewWorld(height, width int) (*World, error) {
	if height <= 0 || width <= 0 {
		return nil, fmt.Errorf("negative height or width")
	}
	cells := make([][]bool, height)
	for i, _ := range cells {
		cells[i] = make([]bool, width)
	}

	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}, nil
}
