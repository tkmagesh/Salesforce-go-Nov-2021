package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//read from file
	file, err := os.Open("data1.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
