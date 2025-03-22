package main

import (
	"fmt"
	"os"
)

func loadAsCalculatorData() todoList {
	file, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("No file, no party :(")
	}
	fmt.Printf("File Contents: %s\n", file)

	return todoList{}
}
