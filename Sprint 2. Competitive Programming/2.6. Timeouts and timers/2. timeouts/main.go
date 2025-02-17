package main

import (
	"errors"
	"fmt"
	"time"
)

func TimeoutFibonacci(n int, timeout time.Duration) (int, error) {
	if n < 0 {
		return 0, errors.New("negative number")
	}
	resultChan := make(chan int)

	timer := time.NewTimer(timeout)

	go func() {
		resultChan <- Fibonacci(n)
	}()

	select {
	case result := <-resultChan:
		return result, nil
	case <-timer.C:
		return 0, errors.New("time's up")
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	res, err := TimeoutFibonacci(5, 3)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(res)
	}
}
