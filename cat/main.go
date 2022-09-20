package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func console_scanner_for_cat() {
	_, e := io.Copy(os.Stdout, os.Stdin)
	if e != nil {
		panic(e)
	}
}

func display(t []byte) {
	for i := 0; i < len(t); i++ {
		z01.PrintRune(rune(t[i]))
	}
}

func main() {
	arguments := os.Args[1:]
	if len(arguments) == 0 {
		console_scanner_for_cat()
	}
	for i := 0; i < len(arguments); i++ {
		data, err := os.ReadFile(arguments[i])
		if err != nil {
			display([]byte("ERROR: " + err.Error() + "\n"))
			os.Exit(1)
		} else {
			display(data)
		}
	}
}
