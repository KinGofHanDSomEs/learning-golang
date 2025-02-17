package main

func CalculateSeriesSum(n int) float64 {
	var sum float64
	for i := 1; i < n+1; i++ {
		sum += 1 / float64(i)
	}
	return sum
}
