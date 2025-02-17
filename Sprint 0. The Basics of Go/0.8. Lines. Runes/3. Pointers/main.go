package main

import "strings"

func isLatin(input string) bool {
	for _, char := range strings.ToLower(input) {
		if char < 97 || char > 122 {
			return false
		}
	}
	return true
}
