package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	t := os.Args
	for i := 1; i < len(t); i++ {
		z := []rune(t[i])
		for j := 0; j < len(z); j++ {
			z01.PrintRune(rune(t[i][j]))
		}
		z01.PrintRune('\n')
	}
}
