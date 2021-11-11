package main

import (
	"fmt"
	"sync"
)

var invocationCount int

var wg sync.WaitGroup = sync.WaitGroup{}

var mutex = sync.Mutex{}

func main() {
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go doSomthing()
	}
	wg.Wait()
	fmt.Println("Invocation count : ", invocationCount)
}

func doSomthing() {
	mutex.Lock()
	invocationCount++
	mutex.Unlock()
	wg.Done()
}
