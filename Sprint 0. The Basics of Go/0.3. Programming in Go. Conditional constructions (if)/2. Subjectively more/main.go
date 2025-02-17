package main

import "fmt"

func main() {
	var number1, number2 int
	fmt.Scanln(&number1, &number2)
	if number1 > number2 {
		fmt.Print("Первое число больше второго")
	} else if number1 < number2 {
		fmt.Print("Второе число больше первого")
	} else {
		fmt.Print("Числа равны")
	}
}
