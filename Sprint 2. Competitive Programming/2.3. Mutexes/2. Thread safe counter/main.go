package main

import (
	"sync"
)

type Сount interface {
	Increment()
	GetValue() int
}

type Counter struct {
	value int
	mu    sync.RWMutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) GetValue() int {
	c.mu.Lock()
	res := c.value
	c.mu.Unlock()
	return res
}
