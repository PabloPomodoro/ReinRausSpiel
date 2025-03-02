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
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Invalid input")
	scanner.Scan()
}

func main() {
	var num1, num2 float64
	var operator string

	for {
		fmt.Print("Enter first number: ")
		_, err := fmt.Scanf("%f", &num1)

		if err != nil {
			handleInvalidInput()
			continue
		}
		break
	}

	for {
		fmt.Print("Enter operator (+, -, *, /): ")
		_, err := fmt.Scanf("%s", &operator)

		if err != nil {
			fmt.Println("Choose operator")
			continue
		}
		break
	}

	for {
		fmt.Print("Enter second number: ")
		_, err := fmt.Scanf("%f", &num2)

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
	default:
		fmt.Println("Invalid operator!")
	}
}
