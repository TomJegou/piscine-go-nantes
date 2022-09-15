package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name_slice := []byte(os.Args[0][2:])
	for i := 0; i < len(name_slice); i++ {
		z01.PrintRune(rune(name_slice[i]))
	}
	z01.PrintRune('\n')
}
