package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fibonacci(ch)
	for no := range ch {
		fmt.Println(no)
	}
}

func fibonacci(ch chan int) {
	done := time.After(10 * time.Second)

	/*
		done := func() chan string {
			ch := make(chan string)
			go func() {
				time.Sleep(10 * time.Second)
				ch <- "done"
			}()
			return ch
		}()
	*/

	x, y := 0, 1

	for {
		select {
		case ch <- x:
			x, y = y, x+y
			time.Sleep(time.Millisecond * 500)
		case <-done:
			fmt.Println("Exiting from fibonacci")
			close(ch)
			return
		}
	}
}
