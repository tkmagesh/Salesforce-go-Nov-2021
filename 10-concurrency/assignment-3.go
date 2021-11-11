/* Modify this program in such a way that it keeps generting the fibonacci numbers until the users hits enter key */

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan bool)

	go fibonacci(ch, done)
	go func() {
		var input string
		fmt.Scanln(&input)
		done <- true
	}()

	for no := range ch {
		fmt.Println(no)
	}
}

func fibonacci(ch chan int, done chan bool) {

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
