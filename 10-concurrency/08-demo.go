package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main started")
	ch := make(chan int)
	go add(100, 200, ch)
	/* schedule some time consuming go routines */
	result := <-ch //channel "read" is a blocking operation
	fmt.Println("result :", result)
}

func add(x, y int, ch chan int) {
	fmt.Println("add started")
	time.Sleep(5 * time.Second)
	result := x + y
	ch <- result
	fmt.Println("add completed")
}

/*
x := <- ch => reading data from the channel
ch <- 10 => writing data in to the channel
*/
