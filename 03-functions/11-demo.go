package main

import (
	"errors"
	"fmt"
)

var DenominatorZeroError = errors.New("denominator cannot be zero")

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered in main", err)
			return
		}
		fmt.Println("successful exit")
	}()
	fmt.Println(divide(100, 0))
}

func divide(x, y int) int {
	if y == 0 {
		panic(DenominatorZeroError)
	}
	return x / y
}
