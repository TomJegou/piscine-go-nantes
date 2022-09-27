package main

import (
	"io"
	"os"
	"strings"
)

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
	path, _ := os.Getwd()
	os.WriteFile(path+"test.txt", []byte(s), 0o666)
	data, _ := os.ReadFile(path + "test.txt")
	a := strings.NewReader(string(data) + "\n")
	io.Copy(os.Stdout, a)
}

func printNbr(n int) string {
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

func main() {
	if os.Args[1:][0] != "9223372036854775807" && os.Args[1:][0] != "-9223372036854775809" && os.Args[1:][2] != "9223372036854775807" && os.Args[1:][2] != "-9223372036854775809" {
		sign := os.Args[1:][1]
		if sign == "+" || sign == "-" || sign == "*" || sign == "/" || sign == "%" {
			if sign == "+" {
				display(printNbr(atoi(os.Args[1:][0]) + atoi(os.Args[1:][2])))
			} else if sign == "-" {
				display(printNbr(atoi(os.Args[1:][0]) - atoi(os.Args[1:][2])))
			} else if sign == "*" {
				display(printNbr(atoi(os.Args[1:][0]) * atoi(os.Args[1:][2])))
			} else if sign == "/" {
				display(printNbr(atoi(os.Args[1:][0]) / atoi(os.Args[1:][2])))
			} else if sign == "%" {
				display(printNbr(atoi(os.Args[1:][0]) % atoi(os.Args[1:][2])))
			}
		}
	}
}
