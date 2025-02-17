package main

import "fmt"

func main() {
	var number1, number2, number3 int
	fmt.Scanln(&number1, &number2, &number3)
	if number1 == 0 || number2 == 0 || number3 == 0 {
		fmt.Print("Некорректный ввод")
		return
	}
	if number1 == number2 && number2 == number3 {
		fmt.Print("Все числа равны")
		return
	}
	if number1 == number2 || number1 == number3 || number2 == number3 {
		fmt.Print("Два числа равны")
		return
	}
	if number1 != number2 && number1 != number3 && number2 != number3 {
		fmt.Print("Все числа разные")
		return
	}
}
