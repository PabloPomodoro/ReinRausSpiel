package main

import (
	"fmt"
	"log"
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

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter first number: ")
	_, err := fmt.Scanf("%f", &num1)
	for err != nil {
		fmt.Println("Invalid input")

		var discard string
		fmt.Scanln(&discard)

		fmt.Print("Enter first number: ")
		_, err = fmt.Scanf("%f", &num1)
	}

	fmt.Print("Enter operator (+, -, *, /): ")
	_, err = fmt.Scanf("%s", &operator)
	for err != nil {
		fmt.Println("Choose operator")

		var discard string
		fmt.Scanln(&discard)

		fmt.Print("Enter operator (+, -, *, /): ")
		_, err = fmt.Scanf("%s", &operator)
	}

	fmt.Print("Enter second number: ")
	_, err = fmt.Scanf("%f", &num2)
	for err != nil {
		fmt.Println("Invalid input")

		var discard string
		fmt.Scanln(&discard)

		fmt.Print("Enter second number: ")
		_, err = fmt.Scanf("%f", &num2)
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
