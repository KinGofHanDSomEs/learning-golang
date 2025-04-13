package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func SumValuesPipeline(filename string) int {
	return Sum(Filter(NumbersGen(filename)))
}

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

func Filter(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
	}()
	return out
}

func Sum(in <-chan int) int {
	var res int
	for n := range in {
		res += n
	}
	return res
}

func main() {
	fmt.Println(SumValuesPipeline("file"))
}
