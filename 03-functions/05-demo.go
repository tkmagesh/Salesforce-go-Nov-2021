//higher order functions
//functions can be passed as arguments to other functions

package main

import "fmt"

func main() {
	x, y := 100, 200

	/*
		fmt.Printf("Processing %d and %d\n", x, y)
		addResult := add(x, y)
		fmt.Println("Result = ", addResult)

		fmt.Printf("Processing %d and %d\n", x, y)
		subtractResult := subtract(x, y)
		fmt.Println("Result = ", subtractResult)
	*/
	process(x, y, add)
	process(x, y, subtract)
}

func process(x, y int, operation func(int, int) int) {
	fmt.Printf("Processing %d and %d\n", x, y)
	result := operation(x, y)
	fmt.Println("Result = ", result)
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}
