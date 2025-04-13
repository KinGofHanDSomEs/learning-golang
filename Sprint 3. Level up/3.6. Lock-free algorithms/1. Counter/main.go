package main

import (
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int64
}

func (c *Counter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *Counter) Decrement() {
	atomic.AddInt64(&c.value, -1)
}

func (c *Counter) GetValue() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	c := Counter{4}
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increment()
		}()

	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Decrement()
		}()
	}s
	wg.Wait()
	print(c.GetValue())
}
