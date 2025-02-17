package main

import "strings"

func IsPalindrome(input string) bool {
	str := ""
	for _, char := range strings.ToLower(input) {
		if string(char) != " " {
			str += string(char)
		}
	}
	reversed_str := ""
	for _, char := range str {
		reversed_str = string(char) + reversed_str
	}
	if reversed_str == str {
		return true
	}
	return false
}
