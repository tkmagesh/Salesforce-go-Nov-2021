package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func (p *Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

type Products []Product

func (products *Products) Format() string {
	var result string
	for _, p := range *products {
		result += p.Format() + "\n"
	}
	return result
}

func (products *Products) IndexOf(product Product) int {
	for i, p := range *products {
		if p == product {
			return i
		}
	}
	return -1
}

func (products *Products) Includes(product Product) bool {
	return products.IndexOf(product) != -1
}

func (products *Products) Filter(predicate func(Product) bool) Products {
	var result Products
	for _, p := range *products {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

func main() {

	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}

	marker := Product{103, "Marker", 50, 20, "Utencil"}
	fmt.Println("IndexOf marker = ", products.IndexOf(marker))

	fmt.Println("Filter")
	var costlyProductPredicate = func(product Product) bool {
		return product.Cost > 1000
	}
	costlyProducts := products.Filter(costlyProductPredicate)
	fmt.Println("Costly products:")
	fmt.Println(costlyProducts.Format())

	var stationaryProductPredicate = func(product Product) bool {
		return product.Category == "Stationary"
	}
	stationaryProducts := products.Filter(stationaryProductPredicate)
	fmt.Println("Stationary products:")
	fmt.Println(stationaryProducts.Format())
}

/*
Write the following functions
1. IndexOf => return the index of the given product in the collection
2. Includes => return true if the given product is present in the collection else return false
3. Filter => return products that matches the given criteria
	Use cases:
		1. filter all costly products (cost > 1000)
		2. filter all products that are stationary
4. All => return true if ALL the products matches the given criteria else return false
	Use cases:
		1. Are all the products in the collection are costly products (cost > 1000) ?
		2. Are all the products in the collection are stationary products? (category == "Stationary")

5. Any => return true if ANY of the products in the collection matches the given criteria else returns false
	Use cases:
		1. Are there ANY costly product (cost > 1000) ?
		2. Are there ANY stationary products? (category == "Stationary")


Convert the above functions to methods
*/
