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

func add(todoList *todoList, itemName string) {

	item := item{
		name:    itemName,
		dueDate: time.Now(),
		done:    false,
	}

	todoList.items = append(todoList.items, item)
}

func check(itemNumber int) {

}

func remove(itemNumber int) {

}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {

	clearScreen()
	fmt.Println("Welcome to the TODO App!")

	todoList := todoList{}

	var action string
	for {
		var emptyInput string
		fmt.Scanln(&emptyInput)
		clearScreen()

		if len(todoList.items) == 0 {
			fmt.Println("No items yet!")
		} else {
			for itemIndex, item := range todoList.items {
				fmt.Printf("%d: %s, %s, %s\n",
					itemIndex+1,
					item.name,
					item.dueDate.Format("2006-01-02"),
					map[bool]string{true: "✓", false: "✘"}[item.done])
			}
		}
		fmt.Println()

		fmt.Println("Choose actions - add, remove, check, exit")
		_, err := fmt.Scanln(&action)

		if err != nil {
			fmt.Println("Invalid action!")
			continue
		}

		switch action {
		case "add":
			var itemName string

			fmt.Print("Choose item name: ")
			_, err := fmt.Scanln(&itemName)

			if err != nil {
				fmt.Println("Invalid item name!")
				continue
			}

			add(&todoList, itemName)
		case "remove":
			fmt.Println("Which item: ")
			itemNumber, err := fmt.Scanln()

			if err != nil {
				fmt.Println("Invalid item!")
				continue
			}

			remove(itemNumber)
		case "check":
			fmt.Println("Which item: ")
			itemNumber, err := fmt.Scanln()

			if err != nil {
				fmt.Println("Invalid item!")
				continue
			}

			check(itemNumber)
		case "exit":
			goto exit
		default:
			fmt.Println("Invalid action!")
			continue
		}
	}

exit:
	fmt.Println("Goodbye!")

	var emptyInput string
	fmt.Scanln(&emptyInput)

	clearScreen()
}
