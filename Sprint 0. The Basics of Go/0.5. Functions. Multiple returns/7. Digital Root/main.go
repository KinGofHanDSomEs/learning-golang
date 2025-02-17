package main

func CalculateDigitalRoot(n int) int {
	if n < 10 {
		return n
	}
	sum := 0
	for {
		if n < 10 {
			sum += n
			break
		}
		sum += n % 10
		n /= 10
	}
	return CalculateDigitalRoot(sum)
}
