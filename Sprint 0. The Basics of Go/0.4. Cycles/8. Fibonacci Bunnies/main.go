package main

import "fmt"

func main() {
	var num1, num2, input_number = 0, 1, 0
	fmt.Scan(&input_number)
	for {
		if input_number == num1 || num1 > input_number {
			num1, num2 = num2, num1
			break
		}
		if input_number == num2 || num2 > input_number {
			num1, num2 = num2, num1+num2
			break
		}
		num1, num2 = num2, num1+num2
	}
	for i := 0; i < 10; i++ {
		fmt.Println(num1)
		num1, num2 = num2, num1+num2
	}
}
