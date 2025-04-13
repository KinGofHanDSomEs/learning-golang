package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func NumbersGen(filename string) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		r, _ := os.Open(filename)
		scn := bufio.NewScanner(r)
		for scn.Scan() {
			if n, err := strconv.Atoi(scn.Text()); err == nil {
				out <- n
			}
		}
	}()
	return out
}

func main() {
	for num := range NumbersGen("file") {
		fmt.Println(num)
	}
}
