package piscine

import "github.com/01-edu/z01"

func switch_int(a int, b int, t []int) {
	c := t[a]
	t[a] = t[b]
	t[b] = c
}

func reverse(slice_to_reverse []int) {
	j := len(slice_to_reverse) - 1
	for i := 0; i < j; i++ {
		switch_int(i, j, slice_to_reverse)
		j--
	}
}

func preprocessing_nbr(nbr int, base int) []int {
	result := []int{}
	for i := 0; nbr > 0; i++ {
		result = append(result, nbr%base)
		nbr /= base
	}
	return result
}

func display(slice []int, str string) {
	for i := 0; i < len(slice); i++ {
		z01.PrintRune(rune(str[slice[i]]))
	}
}

func check_base(str string) bool {
	if len(str) <= 2 {
		return false
	}
	for i := 0; i < len(str)-1; i++ {
		if str[i] == 45 || str[i] == 43 {
			return false
		}
		for j := len(str) - 1; j > i; j-- {
			if str[j] == str[i] || str[j] == 45 || str[j] == 43 {
				return false
			}
		}
	}
	return true
}

func PrintNbrBase(nbr int, base string) {
	if check_base(base) {
		if nbr < 0 {
			z01.PrintRune(45)
			nbr *= -1
		}
		z := preprocessing_nbr(nbr, len(base))
		reverse(z)
		display(z, base)
	} else {
		z01.PrintRune(78)
		z01.PrintRune(86)
	}
}
