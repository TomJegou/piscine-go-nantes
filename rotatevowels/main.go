package main

import (
	"os"

	"github.com/01-edu/z01"
)

func switch_vowels(i1 int, i2 int, t []byte) {
	c := t[i1]
	t[i1] = t[i2]
	t[i2] = c
}

func isVowel(s byte) bool {
	const all_vowels = "aeiouAEIOU"
	for i := 0; i < len(all_vowels); i++ {
		if s == all_vowels[i] {
			return true
		}
	}
	return false
}

func main() {
	arguments := os.Args[1:]
	if len(arguments) >= 1 {
		s := arguments[0]
		for i := 1; i < len(arguments); i++ {
			s += " " + arguments[i]
		}
		s_in_slice_byte := []byte(s)
		start := 0
		end := len(s_in_slice_byte) - 1
		for i := 0; i < len(s_in_slice_byte); i++ {
			for j := start; start < end; j++ {
				start = j
				if isVowel(s_in_slice_byte[j]) {
					start = j
					break
				}
			}
			for k := end; end > start; k-- {
				end = k
				if isVowel(s_in_slice_byte[k]) {
					end = k
					break
				}
			}
			switch_vowels(start, end, s_in_slice_byte)
			if start < end {
				start++
			}
			if end > start {
				end--
			}
		}
		s = string(s_in_slice_byte)
		for i := 0; i < len(s); i++ {
			z01.PrintRune(rune(s[i]))
		}
	}
	z01.PrintRune('\n')
}
