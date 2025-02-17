package life

import (
	"bufio"
	"errors"
	"math/rand"
	"os"
)

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func NewWorld(height, width int) (*World, error) {
	if height < 0 || width < 0 {
		return &World{}, errors.New("negative height or width")
	}
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}, nil
}

func (w *World) LoadState(filename string) error {
	field := [][]bool{}
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	width, f := 0, true
	for fileScanner.Scan() {
		row := []bool{}
		fileRow := fileScanner.Text()
		if f {
			width = len(fileRow)
			f = false
		}
		if width != len(fileRow) {
			return errors.New("different width")
		}
		for _, val := range fileRow {
			if val == '1' {
				row = append(row, true)
				continue
			}
			row = append(row, false)
		}
		field = append(field, row)
	}
	w.Height, w.Width, w.Cells = len(field), len(field[0]), field
	return nil
}

func (w *World) SaveState(filename string) error {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return err
	}
	for i, row := range w.Cells {
		for _, val := range row {
			if val {
				_, err = file.Write([]byte{'1'})
				if err != nil {
					return err
				}
				continue
			}
			_, err = file.Write([]byte{'0'})
			if err != nil {
				return err
			}
		}
		if i != len(w.Cells)-1 {
			_, err = file.Write([]byte{'\n'})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (w *World) String() string {
	resultString, blackSquare, whiteSquare := "", "\xe2\xac\x9b", "\xe2\xac\x9c"
	for _, row := range w.Cells {
		for _, val := range row {
			if val {
				resultString += blackSquare
				continue
			}
			resultString += whiteSquare
		}
		resultString += "\n"
	}
	return resultString
}

func (w *World) Neighbors(x, y int) int {
	countAlive := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if (i != x || j != y) && w.Cells[(j+w.Height)%w.Height][(i+w.Width)%w.Width] {
				countAlive++
			}
		}
	}
	return countAlive
}

func (w *World) Next(x, y int) bool {
	n := w.Neighbors(x, y)
	alive := w.Cells[y][x]
	if n < 4 && n > 1 && alive {
		return true
	}
	if n == 3 && !alive {
		return true
	}
	return false
}

func NextState(oldWorld, newWorld *World) {
	for i := 0; i < oldWorld.Height; i++ {
		for j := 0; j < oldWorld.Width; j++ {
			newWorld.Cells[i][j] = oldWorld.Next(j, i)
		}
	}
}

func (w *World) RandInit(n int) {
	for _, row := range w.Cells {
		for i := range row {
			if rand.Intn(n) == 1 {
				row[i] = true
			}
		}
	}
}
