package main

func FiveSteps(array [5]int) [5]int {
	array[0], array[4] = array[4], array[0]
	array[1], array[3] = array[3], array[1]
	return array
}
