package main

func Tee[T any](done <-chan struct{}, in <-chan T) (<-chan T, <-chan T) {
	out1, out2 := make(chan T), make(chan T)
	go func() {
		defer close(out1)
		defer close(out2)
		for {
			select {
			case <-done:
				return
			case val, ok := <-in:
				if !ok {
					return
				}
				out1 <- val
				out2 <- val
			}
		}
	}()
	return out1, out2
}
