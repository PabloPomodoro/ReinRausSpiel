package main

import "fmt"

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {

	clearScreen()
	printMessageAndWaitForInput("Welcome to the TODO App!")

	todos, err := loadTodos()
	if err != nil {
		fmt.Println("Could not load saved data - create new todos")
		todos = todoList{}
	}

	startTodoList(todos)
}
