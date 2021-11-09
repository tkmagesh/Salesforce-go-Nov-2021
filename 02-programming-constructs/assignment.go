package main

import "fmt"

func main() {
	for {
		var choice int
		fmt.Println("Enter your choice:")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Scanf("%d", &choice)
		if choice == 5 {
			break
		}
		if choice < 1 || choice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		var num1, num2 int
		fmt.Println("Enter two numbers:")
		fmt.Scanf("%d %d", &num1, &num2)
		var result int
		switch choice {
		case 1:
			result = num1 + num2
		case 2:
			result = num1 - num2
		case 3:
			result = num1 * num2
		case 4:
			result = num1 / num2
		}
		fmt.Println("Result: ", result)
	}
}
