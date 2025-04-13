package main

import (
	"context"
	"sync"
)

func FanIn[T any](ctx context.Context, channels ...<-chan T) <-chan T {
	outputCh := make(chan T)
	wg := sync.WaitGroup{}
	for _, ch := range channels {
		wg.Add(1)
		go func(input <-chan T) {
			defer wg.Done()
			for {
				select {
				case data, ok := <-input:
					if !ok {
						return
					}
					outputCh <- data
				case <-ctx.Done():
					return
				}
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(outputCh)
	}()

	return outputCh
}
