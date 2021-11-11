package main

import (
	"fmt"
	"sync"
	"time"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {
	wg.Add(2)
	//wg.Add(1)
	go f1()

	//wg.Add(1)
	go f2()

	wg.Wait()
	fmt.Println("exiting main")
}

func f1() {
	fmt.Println("f1 invoked")
	wg.Done()
}

func f2() {

	fmt.Println("f2 invocation started")
	time.Sleep(5 * time.Second)
	fmt.Println("f2 invocation completed")
	wg.Done()
}
