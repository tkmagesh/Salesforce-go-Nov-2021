//closures

package main

import "fmt"

func main() {
	increment := getCounter()
	fmt.Println(increment()) //=> 1
	fmt.Println(increment()) //=> 2
	fmt.Println(increment()) //=> 3
	fmt.Println(increment()) //=> 4
}

func getCounter() func() int { // Step - 1
	var counter int           //=> closure // Step - 2
	increment := func() int { //Step - 3
		counter++ // Step - 4
		return counter
	}
	return increment // Step - 5
}
