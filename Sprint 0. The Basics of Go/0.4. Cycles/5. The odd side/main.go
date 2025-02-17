package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n < 0 {
		fmt.Print("Некорректный ввод")
		return
	}
	sum := 0
	for i := 1; i < n+1; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	fmt.Print(sum)
}
