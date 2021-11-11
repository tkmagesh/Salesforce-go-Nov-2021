package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	dataCh := make(chan int)
	evenCh := make(chan int)
	oddCh := make(chan int)
	evenSumCh := make(chan int)
	oddSumCh := make(chan int)

	fileWg := &sync.WaitGroup{}
	processWg := &sync.WaitGroup{}

	fileWg.Add(2)
	go Source("data1.dat", dataCh, fileWg)
	go Source("data2.dat", dataCh, fileWg)

	processWg.Add(4)
	go Splitter(dataCh, evenCh, oddCh, processWg)
	go Sum(evenCh, evenSumCh, processWg)
	go Sum(oddCh, oddSumCh, processWg)
	go Merger(evenSumCh, oddSumCh, "result.dat", processWg)

	fileWg.Wait()
	close(dataCh)

	processWg.Wait()
	fmt.Println("Done")
}

func Source(fileName string, ch chan int, wg *sync.WaitGroup) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		file.Close()
		wg.Done()
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		if val, err := strconv.Atoi(str); err == nil {
			ch <- val
		}
	}

}

func Splitter(ch chan int, oddCh chan int, even chan int, wg *sync.WaitGroup) {
	defer func() {
		close(oddCh)
		close(even)
		wg.Done()
	}()
	for val := range ch {
		if val%2 == 0 {
			even <- val
		} else {
			oddCh <- val
		}
	}

}

func Sum(ch chan int, out chan int, wg *sync.WaitGroup) {
	defer func() {
		close(out)
		wg.Done()
	}()
	var sum int
	for val := range ch {
		sum += val
	}
	out <- sum
}

func Merger(evenSum, oddSum chan int, resultFile string, wg *sync.WaitGroup) {
	file, err := os.Create(resultFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		file.Close()
		wg.Done()
	}()
	for i := 0; i < 2; i++ {
		select {
		case val := <-evenSum:
			file.WriteString(fmt.Sprintf("Even Total : %d\n", val))
		case val := <-oddSum:
			file.WriteString(fmt.Sprintf("Odd Total : %d\n", val))
		}
	}

}
