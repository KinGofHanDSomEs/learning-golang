package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)
	switch {
	case number == 0:
		fmt.Print("Число равно нулю")
	case number > 0 && number%2 == 0:
		fmt.Print("Число положительное и четное")
	case number > 0 && number%2 != 0:
		fmt.Print("Число положительное и нечетное")
	case number < 0 && number%2 == 0:
		fmt.Print("Число отрицательное и четное")
	case number < 0 && number%2 != 0:
		fmt.Print("Число отрицательное и нечетное")
	}
}
