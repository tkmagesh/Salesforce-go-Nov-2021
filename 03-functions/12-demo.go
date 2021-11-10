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
	/* fmt.Println(divide(100, 0)) */
	result, err := divideClient(100, 0)
	if err != nil {
		fmt.Println("Error occurred", err)
		return
	}
	fmt.Println("Result is", result)
}

func divideClient(x, y int) (result int, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = er.(error)
		}
	}()
	result = divide(x, y)
	return
}

//3rd party library
func divide(x, y int) int {
	if y == 0 {
		panic(DenominatorZeroError)
	}
	return x / y
}
