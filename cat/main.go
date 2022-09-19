package main

import (
	"fmt"
	"os"
)

func loop_hello_for_cat() {
	keep := true
	for i := 0; keep; i++ {
		fmt.Println("Hello")
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
			fmt.Println("ERROR: abc: No such file or directory")
		}
		fmt.Println(string(data))
	}
}
