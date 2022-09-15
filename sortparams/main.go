package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Split(t []rune) ([]rune, rune, []rune) {
	var index_middle int = len(t) / 2
	pivot := t[index_middle]
	slice := t[:index_middle]
	slice = append(slice, t[index_middle+1:]...)
	slice_before := []rune{}
	slice_after := []rune{}
	for i := 0; i < len(slice); i++ {
		if slice[i] < pivot {
			slice_before = append(slice_before, slice[i])
		} else {
			slice_after = append(slice_after, slice[i])
		}
	}
	return slice_before, pivot, slice_after
}

func Merge(s_b []rune, p rune, s_a []rune) []rune {
	s_b = append(s_b, p)
	s_b = append(s_b, s_a...)
	return s_b
}

func Quick_Sort(t []rune) []rune {
	if len(t) < 1 {
		return t
	}
	s_1, p, s_2 := Split(t)
	a := Merge(Quick_Sort(s_1), p, Quick_Sort(s_2))
	return a
}

func SortRune(table []rune) {
	copy(table, Quick_Sort(table))
}

func main() {
	t := os.Args[1:]
	s := []rune{}
	for i := 0; i < len(t); i++ {
		l := []rune(t[i])
		s = append(s, l...)
	}
	SortRune(s)
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
		z01.PrintRune('\n')
	}
}
