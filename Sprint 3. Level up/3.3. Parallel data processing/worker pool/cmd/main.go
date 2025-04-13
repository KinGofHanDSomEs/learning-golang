package main

import (
	"hello/workerpool"
)

type Work struct {
	value int
}

func (w *Work) Task() {
	println(w.value)
}

func main() {
	pool := workerpool.NewPool(3)
	for i := 0; i < 10; i++ {
		pool.AddTask(&Work{i})
	}
	pool.StopPool()
}
