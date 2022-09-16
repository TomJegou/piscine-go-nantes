package main

import (
	"os"

	"github.com/01-edu/z01"
)

func pow(m int, n int) int {
	if n == 0 {
		return 1
	}

	result := m
	for i := 0; i < n-1; i++ {
		result = result * m
	}
	return result
}

func convert_int_to_string(start int, end int, str string) int {
	aStringChangeable := []byte(str)
	var result int
	for i := start; i >= end; i-- {
		ascii_digit := int(aStringChangeable[i])
		if ascii_digit >= 48 && ascii_digit <= 57 {
			result += (ascii_digit - 48) * pow(10, len(str)-1-i)
		} else {
			return 0
		}
	}
	return result
}

func Atoi(s string) int {
	var result int
	if len(s) >= 1 {
		aStringChangeable := []byte(s)
		if int(aStringChangeable[0]) == 43 || int(aStringChangeable[0]) == 45 {
			result = convert_int_to_string(len(s)-1, 1, s)
		} else {
			result = convert_int_to_string(len(s)-1, 0, s)
		}
		if int(aStringChangeable[0]) == 45 {
			result *= -1
		}
	}
	return result
}

func convertint_to_alpha(t []string, upper bool) {
	for i := 0; i < len(t); i++ {
		if Atoi(t[i]) > 0 && Atoi(t[i]) <= 26 {
			if upper {
				z01.PrintRune(rune(Atoi(t[i])) - 1 + 65)
			} else {
				z01.PrintRune(rune(Atoi(t[i])) - 1 + 97)
			}
		} else {
			z01.PrintRune(32)
		}
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) > 1 {
		upper := false
		start := 1
		if os.Args[1] == "--upper" {
			upper = true
			start = 2
		}
		convertint_to_alpha(os.Args[start:], upper)
	}
}
