package main

import (
	"fmt"
)

func main() {
	go f1()
	f2()

	//time.Sleep(1 * time.Second)
	/*
		var input string
		fmt.Scanln(&input)
	*/
	fmt.Println("exiting main")
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
