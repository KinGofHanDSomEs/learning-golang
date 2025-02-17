package main

import "fmt"

type Chest struct {
	val []int
	wt  []int
}

func Knapsack(chest *Chest, maxWeight int) (int, []int) {
	n := len(chest.val)
	matrix := make([][]int, n+1)
	for i := range matrix {
		matrix[i] = make([]int, maxWeight+1)
	}
	for item := 1; item <= n; item++ {
		for capacity := 1; capacity <= maxWeight; capacity++ {
			maxcostWithoutCurrent := matrix[item-1][capacity]
			maxcostWithCurrent := 0
			weightOfCurrent := chest.wt[item-1]
			if capacity >= weightOfCurrent {
				maxcostWithCurrent = chest.val[item-1]
				remainingCapacity := capacity - weightOfCurrent
				maxcostWithCurrent += matrix[item-1][remainingCapacity]
			}
			matrix[item][capacity] = max(maxcostWithoutCurrent, maxcostWithCurrent)
		}
	}
	selectedItems := []int{}
	capacity := maxWeight
	for item := n; item > 0; item-- {
		if matrix[item][capacity] != matrix[item-1][capacity] {
			selectedItems = append(selectedItems, item-1)
			capacity -= chest.wt[item-1]
		}
	}
	return matrix[n][maxWeight], selectedItems
}

func main() {
	w := 10
	chest := Chest{
		val: []int{100, 400, 300, 500},
		wt:  []int{5, 4, 6, 3},
	}
	res1, res2 := Knapsack(&chest, w)
	fmt.Println(res1, res2)
}
