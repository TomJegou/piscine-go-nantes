package main

import (
	"os"

	"github.com/01-edu/z01"
)

func switch_in_string_slice(i1 int, i2 int, t []string) {
	c := t[i1]
	t[i1] = t[i2]
	t[i2] = c
}

func selection_sort(t []string) {
	n := len(t)
	for i := 0; i < len(t); i++ {
		v := 0
		x := t[v]
		for j := 0; j < n; j++ {
			if t[j] > x {
				x = t[j]
				v = j
			}
		}
		switch_in_string_slice(v, n-1, t)
		n--
	}
}

func main() {
	t := os.Args[1:]
	selection_sort(t)
	for i := 0; i < len(t); i++ {
		a := []rune(t[i])
		for j := 0; j < len(a); j++ {
			z01.PrintRune(a[j])
		}
		z01.PrintRune('\n')
	}
}
