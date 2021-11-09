package main

import "fmt"

func main() {
	//anonymous function
	func() {
		fmt.Println("anonymous function invoked")
	}()

	//anonymous unction with parameters
	func(x, y int) {
		fmt.Println(x + y)
	}(100, 200)

	//anonymous function with return value
	result := func(x, y int) int {
		return x + y
	}(100, 200)
	fmt.Println(result)
}
