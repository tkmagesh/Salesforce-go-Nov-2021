package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("[@main] deferred")
	}()

	fmt.Println("result from  f1() = ", f1())
	fmt.Println("main completed")
}

func f1() (result int) {
	defer func() {
		fmt.Println("[@f1] deferred - 1")
		result = 500
	}()

	defer func() {
		fmt.Println("[@f1] deferred - 2")
	}()
	fmt.Println("f1 invoked")
	f2()
	fmt.Println("f1 completed")
	return 100
}

func f2() {
	defer func() {
		fmt.Println("[@f2] deferred")
	}()
	fmt.Println("f2 invoked")
	fmt.Println("f2 completed")
}
