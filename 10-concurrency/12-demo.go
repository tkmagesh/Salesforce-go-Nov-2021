package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	done := time.After(20 * time.Second)
	for {
		select {
		case <-tick:
			fmt.Print("tick")
		case <-done:
			fmt.Println("done")
			return
		default:
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
