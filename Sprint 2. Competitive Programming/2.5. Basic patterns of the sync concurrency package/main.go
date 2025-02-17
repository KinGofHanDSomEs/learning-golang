package main

import (
	"fmt"
	"sync"
	"time"
)

func listen(name string, data map[string]string, c *sync.Cond) {
	c.L.Lock()
	c.Wait()

	fmt.Printf("[%s] %s\n", name, data["key"])

	c.L.Unlock()
}

func broadcast(name string, data map[string]string, c *sync.Cond) {
	time.Sleep(time.Second)

	c.L.Lock()

	data["key"] = "value"

	fmt.Printf("[%s] данные получены\n", name)

	c.Broadcast()
	c.L.Unlock()
}

func main() {
	var wg sync.WaitGroup
	data := map[string]string{}

	cond := sync.NewCond(&sync.Mutex{})
	wg.Add(3)
	go func() {
		defer wg.Done()
		listen("слушатель 1", data, cond)
	}()
	go func() {
		defer wg.Done()
		listen("слушатель 2", data, cond)
	}()
	go func() {
		defer wg.Done()
		broadcast("источник", data, cond)
	}()
	wg.Wait()
}
