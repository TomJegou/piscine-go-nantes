package main

import (
	"bufio"
	"os"

	"github.com/01-edu/z01"
)

func console_scanner_for_cat() {
	consolescanner := bufio.NewScanner(os.Stdin)
	for consolescanner.Scan() {
		input := consolescanner.Text()
		display([]byte(input))
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
