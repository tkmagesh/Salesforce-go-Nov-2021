//higher order functions
//functions can be passed as arguments to other functions
//functions can be returned from other functions

package main

import "fmt"

func main() {
	x, y := 100, 200

	loggedAdd := getLoggedOperation(add)
	loggedSubtract := getLoggedOperation(subtract)

	loggedAdd(x, y)
	loggedSubtract(x, y)
}

func getLoggedOperation(operation func(int, int) int) func(int, int) {
	return func(x, y int) {
		fmt.Printf("Processing %d and %d\n", x, y)
		result := operation(x, y)
		fmt.Println("Result = ", result)
	}
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}
