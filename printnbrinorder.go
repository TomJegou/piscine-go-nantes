package piscine

import "github.com/01-edu/z01"

func split_int_in_slice(n int) {
	result := []byte{}
	for i := 0; n > 9; i++ {
		result = append(result, 48+byte(n%10))
		z01.PrintRune(rune(result[len(result)-1]))
		n /= 10
	}
	result = append(result, 48+byte(n%10))
	z01.PrintRune(rune(result[len(result)-1]))
}

func PrintNbrInOrder(n int) {
	split_int_in_slice(n)
}
