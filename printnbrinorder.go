package piscine

import "github.com/01-edu/z01"

func Split_for_byte(t []byte) ([]byte, byte, []byte) {
	var index_middle int = len(t) / 2
	pivot := t[index_middle]
	slice := t[:index_middle]
	slice = append(slice, t[index_middle+1:]...)
	slice_before := []byte{}
	slice_after := []byte{}
	for i := 0; i < len(slice); i++ {
		if slice[i] < pivot {
			slice_before = append(slice_before, slice[i])
		} else {
			slice_after = append(slice_after, slice[i])
		}
	}
	return slice_before, pivot, slice_after
}

func Merge_for_byte(s_b []byte, p byte, s_a []byte) []byte {
	s_b = append(s_b, p)
	s_b = append(s_b, s_a...)
	return s_b
}

func Quick_Sort_for_byte(t []byte) []byte {
	if len(t) < 1 {
		return t
	}
	s_1, p, s_2 := Split_for_byte(t)
	a := Merge_for_byte(Quick_Sort_for_byte(s_1), p, Quick_Sort_for_byte(s_2))
	return a
}

func SortByteTable(table []byte) {
	copy(table, Quick_Sort_for_byte(table))
}

func split_int_in_slice(n int) []byte {
	result := []byte{}
	for i := 0; n > 9; i++ {
		result = append(result, 48+byte(n%10))
		n /= 10
	}
	result = append(result, 48+byte(n%10))
	SortByteTable(result)
	return result
}

func PrintNbrInOrder(n int) {
	slice_byte := split_int_in_slice(n)
	for i := 0; i < len(slice_byte); i++ {
		z01.PrintRune(rune(slice_byte[i]))
	}
}
