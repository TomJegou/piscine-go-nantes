package piscine

import (
	"github.com/01-edu/z01"
)

func Power(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if nb == 1 || nb == 0 {
		return nb
	} else if power == 0 {
		return 1
	}
	a := Power(nb, power/2)
	if power%2 == 0 {
		return a * a
	} else {
		return a * a * nb
	}
}

func convert_int_into_slice(n int) []int {
	sA := []int{}
	count := 0
	for i := 0; n/Power(10, i) >= 9; i++ {
		count++
	}
	for i := count; i >= 0; i-- {
		sA = append(sA, n/Power(10, i))
		n -= (n / Power(10, i)) * Power(10, i)
	}
	return sA
}

func convert_slice_to_string(t []int) {
	for i := 0; i < len(t); i++ {
		z01.PrintRune(48 + rune(t[i]))
	}
}

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune(48)
	} else if n > 0 {
		t := convert_int_into_slice(n)
		convert_slice_to_string(t)
	} else {
		z01.PrintRune(45)
		z := n * -1
		t := convert_int_into_slice(z)
		convert_slice_to_string(t)
	}
}
