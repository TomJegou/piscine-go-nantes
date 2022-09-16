package main

import (
	"os"

	"github.com/01-edu/z01"
)

func convertint_to_alpha(t []byte, upper bool) {
	for i := 0; i < len(t); i++ {
		if t[i] >= '0' && t[i] <= '9' {
			if upper {
				z01.PrintRune(rune(t[i] - 49 + 65))
			} else {
				z01.PrintRune(rune(t[i] - 49 + 97))
			}
		} else {
			z01.PrintRune(32)
		}
	}
	z01.PrintRune('\n')
}

func main() {
	upper := false
	start := 1
	arguments := []byte{}
	if os.Args[1] == "--upper" {
		upper = true
		start = 2
	}
	for i := start; i < len(os.Args); i++ {
		t := []byte(os.Args[i])
		arguments = append(arguments, t...)
	}
	convertint_to_alpha(arguments, upper)
}
