package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

type PerishableProduct struct {
	Product
	Id     int
	Expiry string
}

func main() {
	//creating an instance
	grapes := PerishableProduct{Product: Product{201, "Grapes", 5.0, 20, "Fruit"}, Expiry: "2 Days"}
	fmt.Printf("%#v\n", grapes)

	fmt.Println(grapes.Product.Id, grapes.Name, grapes.Cost)

	//use the applyDiscount function to apply 20% discount the 'grapes'
	applyDiscount(&grapes.Product, 0.2)
	fmt.Printf("%#v\n", grapes)
}

func applyDiscount(product *Product, discount float64) {
	product.Cost *= (1 - discount)
}
