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

func add(todos *todoList, itemName string) {

	item := item{
		name:    itemName,
		dueDate: time.Now(),
		done:    false,
	}

	todos.items = append(todos.items, item)
}

func removeAtIndex(slice []item, index int) []item {
	return append(slice[:index], slice[index+1:]...)
}

func remove(todos *todoList, itemNumber int) {
	todos.items = removeAtIndex(todos.items, itemNumber)
}

func check(todos *todoList, itemNumber int) {

	for i := range todos.items {
		if i == itemNumber {
			todos.items[i].done = true
			break
		}
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func showTodoList(todos todoList) {
	if len(todos.items) == 0 {
		fmt.Println("No items yet!")
	} else {
		for i, item := range todos.items {
			fmt.Printf("%d: %s, %s, %s\n",
				i+1,
				item.name,
				item.dueDate.Format("2006-01-02"),
				map[bool]string{true: "✓", false: "✘"}[item.done])
		}
	}
	fmt.Println()
}

func printMessageAndWaitForInput(message string) {
	fmt.Println(message)
	var emptyInput string
	fmt.Scanln(&emptyInput)
}

func main() {

	clearScreen()
	printMessageAndWaitForInput("Welcome to the TODO App!")

	todos := todoList{}

	var action string
	for {
		clearScreen()
		showTodoList(todos)

		fmt.Println("Choose actions - add, remove, check, exit")
		_, err := fmt.Scanln(&action)

		if err != nil {
			printMessageAndWaitForInput("Invalid action!")
			continue
		}

		switch action {
		case "add":
			var itemName string

			fmt.Print("Choose item name: ")
			_, err := fmt.Scanln(&itemName)

			if err != nil {
				printMessageAndWaitForInput("Invalid item name!")
				continue
			}

			add(&todos, itemName)
		case "remove":
			fmt.Print("Which item: ")
			var itemNumber int
			_, err := fmt.Scanln(&itemNumber)

			if err != nil {
				printMessageAndWaitForInput("Invalid item!")
				continue
			}

			remove(&todos, itemNumber-1)
		case "check":
			fmt.Print("Which item: ")
			var itemNumber int
			_, err := fmt.Scanln(&itemNumber)

			if err != nil {
				printMessageAndWaitForInput("Invalid item!")
				continue
			}

			check(&todos, itemNumber-1)
		case "exit":
			printMessageAndWaitForInput("Goodbye!")
			clearScreen()
			return
		default:
			printMessageAndWaitForInput("Invalid action!")
			continue
		}
	}
}
