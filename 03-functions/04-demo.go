//higher order functions
//functions assigned to variables
package main

import "fmt"

func main() {
	var fn func()
	fn = func() {
		fmt.Println("fn is invoked")
	}
	fmt.Printf("Type of fn is %T\n", fn)
	fn()

	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	println(add(1, 2))
}
