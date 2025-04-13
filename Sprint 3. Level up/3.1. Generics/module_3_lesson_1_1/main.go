package main

type MyConstraint interface {
	int | float32 | float64
}

func Sum[T MyConstraint](nums []T) T {
	var res T
	for _, n := range nums {
		res += n
	}
	return res
}
