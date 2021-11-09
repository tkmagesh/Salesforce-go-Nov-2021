package main

import "fmt"

func main() {
	/*
		var msg string
		msg = "Hello World"
	*/

	/*
		var msg string = "Hello World"
	*/

	//type inference

	var msg = "Hello World"

	//msg = 100
	//The following syntax can be used only in the 'function scope' and not in the 'package scope'
	//msg := "Hello World"

	fmt.Println(msg)
	fmt.Printf("msg is of type %T\n", msg)

	//multiple variables
	/*
		var label string
		var x int
		var y int
		x = 100
		y = 200
		label = "Add Result : "
	*/

	/*
		var label string = "Add Result : "
		var x int = 100
		var y int = 200
	*/

	/*
		var (
			label string
			x     int
			y     int
		)
	*/

	/*
		var (
			label string
			x, y  int
		)

		x = 100
		y = 200
		label = "Add Result : "
	*/

	/*
		var (
			label string = "Add Result : "
			x, y  int    = 100, 200
		)
	*/

	/*
		var (
			label = "Add Result : "
			x, y  = 100, 200
		)
	*/

	/*
		var (
			label, x, y = "Add Result : ", 100, 200
		)
	*/

	/*
		var label, x, y = "Add Result : ", 100, 200
	*/

	label, x, y := "Add Result : ", 100, 200

	fmt.Printf("%s %d + %d = %d\n", label, x, y, x+y)

}
