package game

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func NewWorld(height, width int) *World {
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}
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

func (w *World) Seed() {
	for _, row := range w.Cells {
		for i := range row {
			if rand.Intn(10) == 1 {
				row[i] = true
			}
		}
	}
}

func (w *World) Fill(n int) {
	cells := make([][]bool, w.Height)
	for i := range cells {
		cells[i] = make([]bool, w.Width)
	}
	w.Cells = cells
	for i := 0; i < int(float32(w.Height*w.Width)*(float32(n)/100)); i++ {
		for {
			height, width := rand.Intn(w.Height), rand.Intn(w.Width)
			if !w.Cells[height][width] {
				w.Cells[height][width] = true
				break
			}
		}
	}
}

func RunGame() {
	height := 50
	width := 100
	currentWorld := NewWorld(height, width)
	nextWorld := NewWorld(height, width)
	currentWorld.Seed()
	for {
		dataFile, err := ioutil.ReadFile("state")
		if err != nil {
			fmt.Println("error reading of file 'state'")
			return
		}
		stringFile := strings.Split(string(dataFile), " ")
		if stringFile[1] == "relevant" {
			fill, _ := strconv.Atoi(stringFile[0])
			currentWorld.Fill(fill)
			if err := ioutil.WriteFile("state", []byte(fmt.Sprint(fill)+" used"), 0644); err != nil {
				fmt.Println("error writing of data in file 'state'")
				return
			}
		}
		fmt.Println(currentWorld.String())
		currentWorld.SaveState("game_field")
		NextState(currentWorld, nextWorld)
		currentWorld = nextWorld
		time.Sleep(time.Second)
		fmt.Print("\033[H\033[2J")
	}
}
