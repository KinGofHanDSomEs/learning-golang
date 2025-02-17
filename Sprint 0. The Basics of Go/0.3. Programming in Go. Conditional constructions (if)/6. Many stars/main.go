package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)
	switch {
	case number < 0:
		fmt.Print("Некорректный ввод")
	case number < 10:
		fmt.Print("Число меньше 10")
	case 10 <= number && number < 100:
		fmt.Print("Число меньше 100")
	case 100 <= number && number < 1000:
		fmt.Print("Число меньше 1000")
	case number >= 1000:
		fmt.Print("Число больше или равно 1000")
	}
}
