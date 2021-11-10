package main

import "fmt"

func main() {
	var x int = 10

	//value -> address
	var xPtr *int = &x
	fmt.Println(x, xPtr)

	//address -> value (deferencing)
	var val = *xPtr
	fmt.Println(val)

	fmt.Println("Before incrementing, x = ", x)
	increment(&x)
	fmt.Println("After incrementing, x = ", x)

	n1, n2 := 10, 20
	fmt.Println("Before swapping, n1 = ", n1, "n2 = ", n2)
	swap(&n1, &n2)
	fmt.Println("After swapping, n1 = ", n1, "n2 = ", n2)
}

func increment(no *int) {
	*no++
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
