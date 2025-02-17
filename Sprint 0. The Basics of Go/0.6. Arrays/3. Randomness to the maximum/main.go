package main

func FindMinMaxInArray(array [10]int) (int, int) {
	min_number, max_number := array[0], array[0]
	for i := 0; i < len(array); i++ {
		number := array[i]
		if number > max_number {
			max_number = number
			continue
		}
		if number < min_number {
			min_number = number
			continue
		}
	}
	return max_number, min_number
}
