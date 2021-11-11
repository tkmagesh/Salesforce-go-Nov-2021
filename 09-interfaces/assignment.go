package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

/*
	should be able to sort the products by any attribute
	use sort.Sort() for sorting
*/
func main() {

	products := []Product{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}
	fmt.Println("Initial List")
	fmt.Println(products)
}
