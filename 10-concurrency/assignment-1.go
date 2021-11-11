package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	ch1 := make(chan string)
	ch2 := make(chan string)
	wg.Add(2)
	go printString("Hello", ch1, ch2, wg)
	go printString("World", ch2, ch1, wg)
	ch1 <- "start"
	wg.Wait()
}

func printString(s string, in chan string, out chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		<-in
		fmt.Println(s)
		out <- "done"
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
