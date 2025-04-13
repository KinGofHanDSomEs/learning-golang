package main

import (
	"context"
	"fmt"
)

type sequenced interface {
	getSequence() int
}

type Num struct {
	int
}

func (n Num) getSequence() int {
	return n.int
}

func EvenNumbersGen[T sequenced](ctx context.Context, numbers ...T) <-chan T {
	outCh := make(chan T)
	go func() {
		defer close(outCh)
		for _, n := range numbers {
			select {
			case <-ctx.Done():
				return
			default:
				if n.getSequence()%2 == 0 {
					outCh <- n
				}
			}
		}
	}()
	return outCh
}

func main() {
	for n := range EvenNumbersGen(context.Background(), []Num{{1}, {2}, {3}, {4}}...) {
		fmt.Println(n)
	}
}
