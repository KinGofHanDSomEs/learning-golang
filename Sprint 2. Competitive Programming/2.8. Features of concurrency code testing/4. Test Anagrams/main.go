package main

import (
	"strings"
	"sort"
)

func AreAnagrams(str1, str2 string) bool {
    str1 = strings.ToLower(str1)
    str2 = strings.ToLower(str2)
    if len(str1) != len(str2) {
        return false
    }
    r1 := []rune(str1)
    r2 := []rune(str2)
    sort.Slice(r1, func(i, j int) bool {
        return r1[i] < r1[j]
    })
    sort.Slice(r2, func(i, j int) bool {
        return r2[i] < r2[j]
    })
    return string(r1) == string(r2)
}