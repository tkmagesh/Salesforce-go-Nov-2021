package main

import "fmt"

type MyString string

func (m MyString) Print() {
	fmt.Println(m)
}

func main() {
	var s MyString = MyString("This is a string")
	s.Print()
}
