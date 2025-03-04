package main

import (
	"fmt"
	"time"
)

type item struct {
	name    string
	dueDate time.Time
	done    bool
}

type todoList struct {
	name  string
	items []item
}

func main() {

	fmt.Println("Welcome to the TODO App!")

	var action string
	for {
		fmt.Print("\033[H\033[2J")
		// list
		fmt.Println("THE LIST THE LIST THE LIST")

		fmt.Println("Choose actions - add, remove, check, exit")
		_, err := fmt.Scanln(&action)

		if err != nil {
			fmt.Println("Invalid action")
			continue
		}

		switch action {
		case "add":
			// add
		case "remove":
			// remove
		case "check":
			// check
		case "exit":
			goto exit
		default:
			fmt.Println("Invalid action")
			continue
		}
	}

exit:
	fmt.Println("Goodbye!")
}
