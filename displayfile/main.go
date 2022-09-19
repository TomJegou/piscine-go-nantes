package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) == 0 {
		fmt.Println("File name missing")
	} else if len(arguments) > 1 {
		fmt.Println("Too many arguments")
	} else {
		data, _ := os.ReadFile(arguments[0])
		fmt.Print(string(data))
	}
}
