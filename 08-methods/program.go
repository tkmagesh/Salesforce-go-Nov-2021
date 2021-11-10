package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func main() {

	product := Product{Id: 100, Name: "Pen", Cost: 5.0, Units: 10, Category: "Stationery"}
	//print(product)
	product.print()

	fmt.Printf("Before applying discount : \n%#v\n", product)
	product.applyDiscount(0.50)
	fmt.Printf("After applying discount : \n%#v\n", product)

}

func (product Product) print() {
	fmt.Printf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s\n", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func (product *Product) applyDiscount(discount float64) {
	product.Cost *= (1 - discount)
}
