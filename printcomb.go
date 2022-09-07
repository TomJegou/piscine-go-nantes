package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	const all_digit string = "0123456789"
	for left_digit := 0; left_digit <= 7; left_digit++ {
		for middle_digit := left_digit + 1; middle_digit <= 8; middle_digit++ {
			for right_digit := middle_digit + 2; right_digit <= 9; right_digit++ {
				z01.PrintRune(rune(all_digit[left_digit]))
				z01.PrintRune(rune(all_digit[middle_digit]))
				z01.PrintRune(rune(all_digit[right_digit]))
				z01.PrintRune(44)
				z01.PrintRune(32)
			}
		}
	}
}
