package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/01-edu/z01"
)

func switch_in_string_slice(i1 int, i2 int, t []byte) {
	c := t[i1]
	t[i1] = t[i2]
	t[i2] = c
}

func selection_sort(t []byte) {
	end := len(t) - 1
	for i := 0; i < len(t); i++ {
		biggest := int(t[0])
		counter := 0
		for j := 0; j <= end; j++ {
			if int(t[j]) > biggest {
				biggest = int(t[j])
				counter = j
			}
		}
		switch_in_string_slice(counter, end, t)
		end -= 1
	}
}

func display_manual() {
	fmt.Println("--insert")
	z01.PrintRune(32)
	z01.PrintRune(32)
	fmt.Println("-i")
	fmt.Print("\t")
	z01.PrintRune(32)
	fmt.Println("This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	z01.PrintRune(32)
	z01.PrintRune(32)
	fmt.Println("-o")
	fmt.Print("\t")
	z01.PrintRune(32)
	fmt.Println("This flag will behave like a boolean, if it is called it will order the argument.")
}

func _order_(s string) string {
	s_byte := []byte(s)
	selection_sort(s_byte)
	return string(s_byte)
}

func _insert_(s string, s_to_insert string) string {
	return s + s_to_insert
}

func main() {
	arguments := os.Args[1:]
	fmt.Println(arguments)
	if len(arguments) == 0 || arguments[0] == "--help" || arguments[0] == "-h" {
		display_manual()
	} else {
		var s_to_insert string
		order := false
		insert := false
		search_insert := true
		s := arguments[len(arguments)-1]
		for i := 0; i < len(arguments)-1; i++ {
			if arguments[i] == "-o" || arguments[i] == "--order" {
				order = true
			} else if arguments[i] == "-i" {
				insert = true
				s_to_insert = arguments[i+1]
				search_insert = false
			} else if search_insert {
				insert = true
				s_to_insert = strings.Split(arguments[i], "=")[1]
			}
		}
		if order && !insert {
			fmt.Println(_order_(s))
		} else if insert && order {
			z := _insert_(s, s_to_insert)
			fmt.Println(_order_(z))
		} else if insert && !order {
			z := _insert_(s, s_to_insert)
			fmt.Println(z)
		}
	}
}
