package main

import "sort"

func Permutations(input string) []string {
	var res []string
	permute([]rune(input), 0, &res)
	sort.Strings(res)
	return res
}

func permute(strs []rune, i int, res *[]string) {
	if i == len(strs)-1 {
		*res = append(*res, string(strs))
		return
	}
	seen := make(map[rune]bool)
	for j := i; j < len(strs); j++ {
		if seen[strs[j]] {
			continue
		}
		seen[strs[j]] = true
		strs[i], strs[j] = strs[j], strs[i]
		permute(strs, i+1, res)
		strs[i], strs[j] = strs[j], strs[i]
	}
}