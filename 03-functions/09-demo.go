//handling errors
package main

import (
	"errors"
	"fmt"
)

var DenominatorZeroError = errors.New("denominator cannot be zero")

func main() {
	//fmt.Println(divide(10, 0))
	result, err := divide(10, 0)
	if err == DenominatorZeroError {
		fmt.Println("Make sure that the denominator is not zero")
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, DenominatorZeroError
	}
	return x / y, nil
}
