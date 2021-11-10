package main

import (
	"fmt"
)

var operations = map[int]func(float32, float32) float32{
	1: add,
	2: subtract,
	3: multiply,
	4: divide,
}

func main() {
	for {
		choice := getUserChoice()
		if choice == 5 {
			break
		}
		if choice < 1 || choice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		num1, num2 := getOperands()
		result := doCalculate(num1, num2, choice)
		fmt.Println("Result: ", result)
	}

}

func doCalculate(num1, num2 float32, choice int) float32 {
	var result float32
	result = operations[choice](num1, num2)
	return result
}

func getOperands() (float32, float32) {
	var num1, num2 float32
	fmt.Println("Enter two numbers:")
	fmt.Scanf("%f %f", &num1, &num2)
	return num1, num2
}

func getUserChoice() int {
	var choice int
	fmt.Println("Enter your choice:")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Scanf("%d", &choice)
	return choice
}

func add(x, y float32) float32 {
	return x + y
}

func subtract(x, y float32) float32 {
	return x - y
}

func multiply(x, y float32) float32 {
	return x * y
}

func divide(x, y float32) float32 {
	return x / y
}
