package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i < n+1; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
