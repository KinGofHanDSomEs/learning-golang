package main

func DoubleNumbers(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case n, ok := <-in:
				if !ok {
					return
				}
				out <- n * 2
			}
		}
	}()
	return out
}
