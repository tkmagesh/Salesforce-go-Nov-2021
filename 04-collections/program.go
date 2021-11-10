package main

import "fmt"

func main() {
	//Array
	fmt.Println("Array")

	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	//var nos [5]int = [...]int{3, 1, 4, 2, 5}
	var nos = [5]int{3, 1, 4, 2, 5}

	fmt.Println(nos)

	fmt.Println("Iterating using an indexer")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Println(nos[idx])
	}

	fmt.Println("Iterating using range")
	/*
		for idx, value := range nos {
			fmt.Println(value, idx)
		}
	*/

	for _, value := range nos {
		fmt.Println(value)
	}
}
