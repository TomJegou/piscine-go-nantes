package piscine

import "github.com/01-edu/z01"

func PrintNbrBase_forconvertbase(nbr int, base string) string {
	result := ""
	if check_base(base) {
		if nbr < 0 {
			result += "-"
		} else if nbr == 0 {
			z := []byte{0}
			for i := 0; i < len(z); i++ {
				result += string(base[z[i]])
			}
		}
		z := preprocessing_nbr(nbr, len(base))
		reverse_slice_byte_for_atoibase(z)
		for i := 0; i < len(z); i++ {
			result += string(base[z[i]])
		}
	} else {
		z01.PrintRune(78)
		z01.PrintRune(86)
	}
	return result
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	number := AtoiBase(nbr, baseFrom)
	return PrintNbrBase_forconvertbase(number, baseTo)
}
