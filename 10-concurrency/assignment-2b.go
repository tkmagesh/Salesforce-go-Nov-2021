package main

import "fmt"

func main() {
	data := []int{4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55}
	dataSet1 := data[:len(data)/2]
	dataSet2 := data[len(data)/2:]
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch1 <- sum(dataSet1...)
	}()
	go func() {
		ch2 <- sum(dataSet2...)
	}()
	result := <-ch1 + <-ch2
	fmt.Println("Sum = ", result)
}

func sum(nos ...int) int {
	var result int
	for _, v := range nos {
		result += v
	}
	return result
}

/*
	Divide & Conquer
	Break the dataset into 2 sets, add them using 2 goroutines and print the final result
*/
