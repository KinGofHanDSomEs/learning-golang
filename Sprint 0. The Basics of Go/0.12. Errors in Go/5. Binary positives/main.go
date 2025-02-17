package main

import (
	"fmt"
)

func IntToBinary(num int) (string, error) {
	if num < 0 {
		return "", fmt.Errorf("negative numbers are not allowed")
	}
	res := ""
	for {
		if num <= 0 {
			return res, nil
		}
		res = fmt.Sprintf("%d", num%2) + res
		num /= 2
	}
}
