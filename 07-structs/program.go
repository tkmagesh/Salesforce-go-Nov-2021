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
	/*
		product := Product{}
		product.Id = 100
		product.Name = "Pen"
		product.Cost = 5.0
		product.Units = 10
		product.Category = "Stationery"
	*/

	//product := Product{100, "Pen", 5.0, 10, "Stationary"}

	product := Product{Id: 100, Name: "Pen", Cost: 5.0}

	fmt.Printf("%#v\n", product)
	/*
		//productPtr := &Product{}
		productPtr := new(Product)
		fmt.Printf("%#v\n", productPtr)
	*/

	fmt.Printf("Before applying discount : \n%#v\n", product)
	applyDiscount(&product, 0.50)
	fmt.Printf("After applying discount : \n%#v\n", product)

}

func applyDiscount(product *Product, discount float64) {
	product.Cost *= (1 - discount)
}
