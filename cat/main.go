package main

import (
	"log"
	"os"

	"github.com/01-edu/z01"
)

func loop_hello_for_cat() {
	message := "Hello"
	keep := true
	for i := 0; keep; i++ {
		for j := 0; j < len(message); j++ {
			z01.PrintRune(rune(message[j]))
		}
		z01.PrintRune('\n')
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
		loop_hello_for_cat()
	}
	for i := 0; i < len(arguments); i++ {
		data, err := os.ReadFile(arguments[i])
		if err != nil {
			log.Fatal(err)
		} else {
			display(data)
		}
	}
}
