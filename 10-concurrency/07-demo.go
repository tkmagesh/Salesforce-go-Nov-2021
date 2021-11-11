package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup = sync.WaitGroup{}

func main() {
	fmt.Println("main started")
	ch := make(chan int)
	wg.Add(1)
	go add(100, 200, ch)
	result := <-ch
	wg.Wait()
	fmt.Println("result :", result)
}

func add(x, y int, ch chan int) {
	defer wg.Done()
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
