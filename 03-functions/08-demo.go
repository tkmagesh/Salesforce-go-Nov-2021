//closures

package main

import "fmt"

func main() {
	increment, decrement := getCounters()

	fmt.Println("Incrementing")
	fmt.Println(increment()) //=> 1
	fmt.Println(increment()) //=> 2
	fmt.Println(increment()) //=> 3
	fmt.Println(increment()) //=> 4

	fmt.Println("Decrementing")
	fmt.Println(decrement()) //=> 3
	fmt.Println(decrement()) //=> 2
	fmt.Println(decrement()) //=> 1
	fmt.Println(decrement()) //=> 0
	fmt.Println(decrement()) //=> -1
}

func getCounters() (func() int, func() int) { // Step - 1
	var counter int           //=> closure // Step - 2
	increment := func() int { //Step - 3
		counter++ // Step - 4
		return counter
	}
	decrement := func() int { //Step - 3
		counter-- // Step - 4
		return counter
	}
	return increment, decrement // Step - 5
}
