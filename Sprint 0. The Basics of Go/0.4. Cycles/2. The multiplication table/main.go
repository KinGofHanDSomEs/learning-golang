package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i < 11; i++ {
		fmt.Println(n, "*", i, "=", n*i)
	}
}
