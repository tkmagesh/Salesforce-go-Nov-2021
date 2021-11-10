package main

import (
	"fmt"
	"modularity-demo/calculator"
	numUtils "modularity-demo/calculator/utils"

	"github.com/fatih/color"
)

func main() {
	fmt.Println(calculator.Add(10, 20))
	fmt.Println(calculator.Subtract(10, 20))
	/* fmt.Println(calculator.multiply(10, 20)) */
	color.Yellow("Op count = %d\n", calculator.GetOpCount())
	//fmt.Printf("Op count = %d\n", calculator.GetOpCount())
	fmt.Println(numUtils.IsEven(201))
}
