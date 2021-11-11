package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

//Implement fmt.Stringer
func (p Product) String() string {
	return fmt.Sprintf("Id: %d, Name: %s, Cost: %f, Units: %d, Category: %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

type Products []Product

//Implement fmt.Stringer
func (p Products) String() string {
	result := ""
	for _, v := range p {
		result += fmt.Sprintf("%s\n", v)
	}
	return result
}

//Implement sort.Interface

func (p Products) Len() int { return len(p) }

func (p Products) Less(i, j int) bool {
	return p[i].Id < p[j].Id
}

func (p Products) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//To sort by name
type ByName struct {
	Products
}

func (byName ByName) Less(i, j int) bool {
	return byName.Products[i].Name < byName.Products[j].Name
}

//To sort by cost
type ByCost struct {
	Products
}

func (byCost ByCost) Less(i, j int) bool {
	return byCost.Products[i].Cost < byCost.Products[j].Cost
}

//Sort helper method
func (products Products) Sort(attrName string) {
	switch attrName {
	case "Id":
		sort.Sort(products)
	case "Name":
		sort.Sort(ByName{products})
	case "Cost":
		sort.Sort(ByCost{products})
	}
}

/*
	should be able to sort the products by any attribute
	use sort.Sort() for sorting
*/
func main() {

	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}
	fmt.Println("Initial List")
	fmt.Println(products)

	//sort by Id
	fmt.Println("Sort by Id")
	//sort.Sort(products)
	products.Sort("Id")
	fmt.Println(products)

	//sort by Name
	fmt.Println("Sort by Name")
	//sort.Sort(ByName{products})
	products.Sort("Name")
	fmt.Println(products)

	//sort by Cost
	fmt.Println("Sort by Cost")
	//sort.Sort(ByCost{products})
	products.Sort("Cost")
	fmt.Println(products)
}
