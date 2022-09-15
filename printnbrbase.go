package piscine

import "github.com/01-edu/z01"

func switch_str_forprintnbrbase(a int, b int, t []byte) {
	c := t[a]
	t[a] = t[b]
	t[b] = c
}

func reverse_slice_byte_forprintnbrbase(t []byte) {
	j := len(t) - 1
	for i := 0; i < j; i++ {
		switch_str_forprintnbrbase(i, j, t)
		j--
	}
}

func preprocessing_nbr(nbr int, base int) []byte {
	result := []byte{}
	for i := 0; nbr != 0; i++ {
		a := nbr % base
		if a < 0 {
			a *= -1
		}
		result = append(result, byte(a))
		nbr /= base
	}
	return result
}

func display(t []byte, base string) {
	for i := 0; i < len(t); i++ {
		z01.PrintRune(rune(base[t[i]]))
	}
}

func check_base(str string) bool {
	if len(str) < 2 {
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
		} else if nbr == 0 {
			z := []byte{0}
			display(z, base)
		}
		z := preprocessing_nbr(nbr, len(base))
		reverse_slice_byte_forprintnbrbase(z)
		display(z, base)
	} else {
		z01.PrintRune(78)
		z01.PrintRune(86)
	}
}
