//package declaration
package main

//import dependency packages
import "fmt"

//package level variables

//package init function

//main function
func main() {
	var msg string
	msg = "Welcome to Golang!!"
	//fmt.Println("Hello World!")
	fmt.Printf("Hello World! %s\n", msg)

	//accepting input from user
	var name string
	fmt.Println("Enter the name :")
	fmt.Scanf("%s\n", &name)
	fmt.Printf("Hello %s\n", name)

	var no int
	fmt.Println("Enter the number :")
	fmt.Scanf("%d\n", &no)
	fmt.Printf("%d\n", no)
}

//other functions

/*
	To run the program
		go run hello-world.go
*/
