package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type SafeCounter struct {
	counter int32
}

func (c *SafeCounter) Increment() {
	atomic.AddInt32(&c.counter, 1)
}

var wg sync.WaitGroup = sync.WaitGroup{}

var sc = &SafeCounter{}

func main() {
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go doSomthing()
	}
	wg.Wait()
	fmt.Println("Invocation count : ", sc.counter)
}

func doSomthing() {
	sc.Increment()
	wg.Done()
}
