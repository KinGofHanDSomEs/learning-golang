package main

import "fmt"

func main() {
	var number1, number2 int
	fmt.Scanln(&number1, &number2)
	if number1 == 0 || number2 == 0 {
		fmt.Print("Одно из чисел равно нулю")
		return
	}
	if number1 > 0 && number2 > 0 {
		fmt.Print("Оба числа положительные")
		return
	}
	if number1 < 0 && number2 < 0 {
		fmt.Print("Оба числа отрицательные")
		return
	}
	fmt.Print("Одно число положительное, а другое отрицательное")
}
