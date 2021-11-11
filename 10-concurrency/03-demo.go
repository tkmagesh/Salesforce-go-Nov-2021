package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg *sync.WaitGroup = &sync.WaitGroup{}
	wg.Add(2)
	//wg.Add(1)
	go f1(wg)

	//wg.Add(1)
	go f2(wg)

	wg.Wait()
	fmt.Println("exiting main")
}

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 invoked")
	wg.Done()
}

func f2(wg *sync.WaitGroup) {

	fmt.Println("f2 invocation started")
	time.Sleep(5 * time.Second)
	fmt.Println("f2 invocation completed")
	wg.Done()
}
