package main

import "fmt"

func ToString[T any](done <-chan struct{}, valueStream <-chan T) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case value, ok := <-valueStream:
				if !ok {
					return
				}
				out <- fmt.Sprint(value)

			}
		}
	}()
	return out
}
