package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := 0; i < len(os.Args[0][2:]); i++ {
		z01.PrintRune(rune(os.Args[0][2:][i]))
	}
	z01.PrintRune('\n')
}
