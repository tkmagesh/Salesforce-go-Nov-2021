package main

import "fmt"

type Entity struct {
}

func main() {
	var x interface{}
	x = "hello"
	x = 10
	x = true
	x = nil
	x = []int{1, 2, 3}

	//x = 1000
	if no, success := x.(int); success {
		fmt.Println(no * 2)
	} else {
		fmt.Println("x is not an int")
	}

	//x = true
	x = Entity{}
	switch val := x.(type) {
	case int:
		fmt.Println(val * 2)
	case string:
		fmt.Println(val + " world")
	case bool:
		fmt.Println("Boolean = ", val)
	case nil:
		fmt.Println("nil")
	case []int:
		fmt.Println("list of integers")
	case Entity:
		fmt.Println("Entity")
	default:
		fmt.Println("unknown type")
	}
}
