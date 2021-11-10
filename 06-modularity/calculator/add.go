package calculator

import "fmt"

func Add(x, y int) int {
	opCount++
	return x + y
}

func DoMultiply() {
	fmt.Println(multiply(100, 200))
}
