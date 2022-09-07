package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	const ALL_DIGIT string = "0123456789"
	for left_left_digit := 0; left_left_digit <= 9; left_left_digit++ {
		for left_right_digit := 0; left_right_digit <= 9; left_right_digit++ {
			for right_left_digit := left_left_digit; right_left_digit <= 9; right_left_digit++ {
				for right_right_digit := left_right_digit + 1; right_right_digit <= 9; right_right_digit++ {
					if left_left_digit == 9 && left_right_digit == 8 && right_left_digit == 9 && right_right_digit == 9 {
						z01.PrintRune(rune(ALL_DIGIT[left_left_digit]))
						z01.PrintRune(rune(ALL_DIGIT[left_right_digit]))
						z01.PrintRune(32)
						z01.PrintRune(rune(ALL_DIGIT[right_left_digit]))
						z01.PrintRune(rune(ALL_DIGIT[right_right_digit]))
						z01.PrintRune('\n')
						break
					} else {
						z01.PrintRune(rune(ALL_DIGIT[left_left_digit]))
						z01.PrintRune(rune(ALL_DIGIT[left_right_digit]))
						z01.PrintRune(32)
						z01.PrintRune(rune(ALL_DIGIT[right_left_digit]))
						z01.PrintRune(rune(ALL_DIGIT[right_right_digit]))
						z01.PrintRune(44)
						z01.PrintRune(32)
					}
				}
			}
		}
	}
}
