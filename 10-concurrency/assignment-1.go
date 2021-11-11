package main

import "fmt"

func main() {
	printString("Hello")
	printString("World")
}

func printString(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

/* 
Expected Output:
Hello
World
Hello
World
Hello
World
Hello
World
Hello
World

Note : DONOT move the for loop outside the printString function.
*/
*/