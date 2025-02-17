package main

import (
	"sort"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	rune1, rune2 := []rune(strings.ToLower(str1)), []rune(strings.ToLower(str2))
	sort.Slice(rune1, func(i, j int) bool { return rune1[i] < rune1[j] })
	sort.Slice(rune2, func(i, j int) bool { return rune2[i] < rune2[j] })
	for i := 0; i < len(rune1); i++ {
		if rune1[i] != rune2[i] {
			return false
		}
	}
	return true
}
