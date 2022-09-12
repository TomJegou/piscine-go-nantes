package piscine

import (
	"github.com/01-edu/z01"
)

func convert_int_to_string(n int) {
	var all_digit string = "0123456789"
	var count int = 0
	var count_hundred int = 0
	var count_decade int = 0
	var count_unit int = 0
	for hundred := 0; hundred < 10; hundred++ {
		count_hundred = hundred
		for decade := 0; decade < 10; decade++ {
			count_decade = decade
			for unit := 0; unit < 10; unit++ {
				count_unit = unit
				count++
				if count == n {
					count_unit++
					break
				}
			}
			if count == n {
				break
			}
		}
		if count == n {
			break
		}
	}
	if count_hundred != 0 {
		z01.PrintRune(rune(all_digit[count_hundred]))
		z01.PrintRune(rune(all_digit[count_decade]))
		z01.PrintRune(rune(all_digit[count_unit]))
	} else if count_decade != 0 {
		z01.PrintRune(rune(all_digit[count_decade]))
		z01.PrintRune(rune(all_digit[count_unit]))
	} else {
		z01.PrintRune(rune(all_digit[count_unit]))
	}
}
func PrintNbr(n int) {
	var all_digit string = "0123456789"
	if n == 0 {
		z01.PrintRune(rune(all_digit[0]))
	} else if n > 0 {
		convert_int_to_string(n)
	} else {
		z01.PrintRune(45)
		z := n * -1
		convert_int_to_string(z)
	}
}
