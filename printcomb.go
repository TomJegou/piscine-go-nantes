package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	const ALL_DIGIT string = "0123456789"
	for left_digit := 0; left_digit <= 7; left_digit++ {
		for middle_digit := left_digit + 1; middle_digit <= 8; middle_digit++ {
			for right_digit := middle_digit + 1; right_digit <= 9; right_digit++ {
				if left_digit == 7 && middle_digit == 8 && right_digit == 9 {
					z01.PrintRune(rune(ALL_DIGIT[left_digit]))
					z01.PrintRune(rune(ALL_DIGIT[middle_digit]))
					z01.PrintRune(rune(ALL_DIGIT[right_digit]))
					z01.PrintRune('\n')
					break
				} else {
					z01.PrintRune(rune(ALL_DIGIT[left_digit]))
					z01.PrintRune(rune(ALL_DIGIT[middle_digit]))
					z01.PrintRune(rune(ALL_DIGIT[right_digit]))
					z01.PrintRune(44)
					z01.PrintRune(32)
				}
			}
		}
	}
}
