package main

import "strings"

func CountingSort(contacts []string) map[string]int {

	res_map := make(map[string]int)
	for _, char := range contacts {
		res_map[string(strings.ToLower(char))]++
	}
	return res_map
}
