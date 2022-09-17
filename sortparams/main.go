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
	end := len(t) - 1
	for i := 0; i <= end; i++ {
		biggest := t[0]
		counter := 0
		for j := 0; j <= end; j++ {
			if t[j] > biggest {
				biggest = t[j]
				counter = j
			}
		}
		switch_in_string_slice(counter, end, t)
		end -= 1
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
