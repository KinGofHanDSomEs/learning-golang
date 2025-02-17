package main

import "fmt"

func SortNames(names []string) {
	for {
		k := 0
		for i := 0; i < len(names)-1; i++ {
			if names[i] > names[i+1] {
				names[i], names[i+1] = names[i+1], names[i]
				k++
			}
		}
		if k == 0 {
			break
		}
	}
}

func main() {
	array := []string{"Варвара", "Есения", "Арина", "Аксинья"}
	SortNames(array)
	fmt.Println(array)
}
