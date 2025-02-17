package main

import "strings"

func ConcatenateStrings(str1, str2 string) string {
	return strings.Join([]string{str1, str2}, " ")
}
