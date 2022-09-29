package main

import (
	"os"
)

func validity_number(s string) bool {
	slice := []byte(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] < 48 || slice[i] > 57 {
			if slice[i] != 45 {
				return false
			}
		}
	}
	return true
}

func validity_sign(s string) bool {
	return s == "-" || s == "+" || s == "*" || s == "/" || s == "%"
}

func pow(n int, m int) int {
	if n == 1 || m == 1 {
		return n
	}
	if m == 0 {
		return 1
	}
	a := pow(n, m/2)
	if n%2 == 0 {
		return a * a
	} else {
		return a * a * n
	}
}

func atoi(s string) int {
	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == "-" {
			result *= -1
		} else {
			result += int((byte(s[i]) - 48)) * pow(10, len(s)-1-i)
		}
	}
	return result
}

func display(s string) {
	os.Stdout.Write([]byte(s))
}

func printNbr(n int) string {
	if n == 0 {
		return "0"
	}
	all_figit := "0123456789"
	result := ""
	negative := false
	if n < 0 {
		negative = true
		n *= -1
	}
	for n > 0 {
		result = string(all_figit[n%10]) + result
		n /= 10
	}
	if negative {
		result = "-" + result
	}
	return result
}

func overflow_validity(n1 int, n2 int) bool {
	result := true
	if n1 < 0 && n2 < 0 {
		smallest := n1
		if n2 < n1 {
			smallest := n2
		a := -9223372036854775808 - smallest
	}
	return result
}

func main() {
	if len(os.Args[1:]) == 3 {
		if validity_number(os.Args[1:][0]) && validity_number(os.Args[1:][2]) && validity_sign(os.Args[1:][1]) {
			first_number := atoi(os.Args[1:][0])
			second_number := atoi(os.Args[1:][2])
			sign := os.Args[1:][1]
			if overflow_validity(first_number, second_number) {

			}
		}
	}
}
