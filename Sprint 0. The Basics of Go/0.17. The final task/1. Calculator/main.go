package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")
	return ReviewExpression(&expression)
}

func ReviewExpression(expression *string) (float64, error) {
	result, err := ReviewComponent(expression)
	if err != nil {
		return 0, err
	}
	for len(*expression) > 0 {
		switch (*expression)[0] {
		case '+':
			*expression = (*expression)[1:]
			nextComp, err := ReviewComponent(expression)
			if err != nil {
				return 0, err
			}
			result += nextComp
		case '-':
			*expression = (*expression)[1:]
			nextComp, err := ReviewComponent(expression)
			if err != nil {
				return 0, err
			}
			result -= nextComp
		default:
			return result, nil
		}
	}
	return result, nil
}

func ReviewComponent(expression *string) (float64, error) {
	result, err := ReviewFactor(expression)
	if err != nil {
		return 0, err
	}
	for len(*expression) > 0 {
		switch (*expression)[0] {
		case '*':
			*expression = (*expression)[1:]
			nextFactor, err := ReviewFactor(expression)
			if err != nil {
				return 0, err
			}
			result *= nextFactor
		case '/':
			*expression = (*expression)[1:]
			nextFactor, err := ReviewFactor(expression)
			if err != nil {
				return 0, err
			}
			if nextFactor == 0 {
				return 0, fmt.Errorf("деление на ноль")
			}
			result /= nextFactor
		default:
			return result, nil
		}
	}
	return result, nil
}

func ReviewFactor(expression *string) (float64, error) {
	if len(*expression) == 0 {
		return 0, fmt.Errorf("неожиданный конец выражения")
	}
	if (*expression)[0] == '(' {
		*expression = (*expression)[1:]
		result, err := ReviewExpression(expression)
		if err != nil {
			return 0, err
		}
		if len(*expression) == 0 || (*expression)[0] != ')' {
			return 0, fmt.Errorf("пропущены круглые скобки")
		}
		*expression = (*expression)[1:]
		return result, nil
	}
	return ReviewNumber(expression)
}

func ReviewNumber(expression *string) (float64, error) {
	i := 0
	for i < len(*expression) && (unicode.IsDigit(rune((*expression)[i])) || (*expression)[i] == '.') {
		i++
	}
	if i == 0 {
		return 0, fmt.Errorf("неверное число, найдено '%c'", (*expression)[0])
	}
	num, err := strconv.ParseFloat((*expression)[:i], 64)
	if err != nil {
		return 0, fmt.Errorf("неверное число, ошибка: %v", err)
	}
	*expression = (*expression)[i:]
	return num, nil
}
