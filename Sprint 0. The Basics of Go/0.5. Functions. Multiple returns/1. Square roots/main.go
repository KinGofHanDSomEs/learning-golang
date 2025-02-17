package main

import (
	"fmt"
	"math"
)

func SqRoots() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	discriminant := math.Pow(b, 2) - 4*a*c
	if discriminant == 0 {
		fmt.Print(-b / (2 * a))
	} else if discriminant > 0 {
		fmt.Print((-b-math.Sqrt(discriminant))/(2*a), (-b+math.Sqrt(discriminant))/(2*a))
	} else {
		fmt.Print(0, 0)
	}
}
