package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	counter int
	sync.Mutex
}

func (c *SafeCounter) Increment() {
	c.Lock()
	{
		c.counter++
	}
	c.Unlock()
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
