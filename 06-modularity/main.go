package main

import (
	"fmt"
	"modularity-demo/calculator"
	numUtils "modularity-demo/calculator/utils"
)

func main() {
	fmt.Println(calculator.Add(10, 20))
	fmt.Println(calculator.Subtract(10, 20))
	/* fmt.Println(calculator.multiply(10, 20)) */
	fmt.Println("Op count =", calculator.GetOpCount())
	fmt.Println(numUtils.IsEven(201))
}
