package main

import "fmt"

func IsPowerOfTwoRecursive(n int) {
	for {
		if n == 1 {
			fmt.Print("YES")
			break
		}
		if n%2 != 0 {
			fmt.Print("NO")
			break
		}
		if n == 2 {
			fmt.Print("YES")
			break
		}
		n /= 2
	}
}
