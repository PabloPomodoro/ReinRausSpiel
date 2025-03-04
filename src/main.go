package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func add(numbers ...float64) {

	fmt.Printf(
		"%.2f + %.2f = %.2f\n",
		numbers[0], numbers[1], numbers[0]+numbers[1])
}

func subtract(numbers ...float64) {

	fmt.Printf(
		"%.2f - %.2f = %.2f\n",
		numbers[0], numbers[1], numbers[0]-numbers[1])
}

func multiply(numbers ...float64) {

	fmt.Printf(
		"%.2f * %.2f = %.2f\n",
		numbers[0], numbers[1], numbers[0]*numbers[1])
}

func divide(numbers ...float64) {

	fmt.Printf(
		"%.2f / %.2f = %.2f\n",
		numbers[0], numbers[1], numbers[0]/numbers[1])
}

func handleInvalidInput() {
	fmt.Println("Invalid input")
	bufio.NewScanner(os.Stdin).Scan()
}

func calculator() int32 {
	var num1, num2 float64
	var operator string

	for {
		fmt.Print("Enter first number: ")
		_, err := fmt.Scanln(&num1)

		if err != nil {
			handleInvalidInput()
			continue
		}
		break
	}

	for {
		fmt.Print("Enter operator (+, -, *, /): ")
		_, err := fmt.Scanln(&operator)

		if err != nil {
			fmt.Println("Choose operator")
			continue
		}
		break
	}

	for {
		fmt.Print("Enter second number: ")
		_, err := fmt.Scanln(&num2)

		if err != nil {
			handleInvalidInput()
			continue
		}
		break
	}

	switch operator {
	case "+":
		add(num1, num2)
	case "-":
		subtract(num1, num2)
	case "*":
		multiply(num1, num2)
	case "/":
		if num2 == 0 {
			log.Fatal("Error: Division by zero is not allowed.")
		}
		divide(num1, num2)
	case "#":
		return -1
	default:
		fmt.Println("Invalid operator!")
	}

	fmt.Println("Press Enter to continue...")
	fmt.Scanln()

	return 0
}

func main() {

	for {
		if calculator() == -1 {
			break
		}
		fmt.Print("\033[H\033[2J")
	}

	fmt.Println("GOOD BYE")
}
