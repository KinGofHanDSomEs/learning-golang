package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fact := 1
	for i := 1; i < n+1; i++ {
		fact *= i
	}
	fmt.Print(fact)
}
