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

func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s\n", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

type By func(p1, p2 *Product) bool

func (by By) Sort(products []Product) {
	ps := &productSorter{
		products: products,
		by:       by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// productSorter joins a By function and a slice of Products to be sorted.
type productSorter struct {
	products []Product
	by       func(p1, p2 *Product) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *productSorter) Len() int {
	return len(s.products)
}

// Swap is part of sort.Interface.
func (s *productSorter) Swap(i, j int) {
	s.products[i], s.products[j] = s.products[j], s.products[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *productSorter) Less(i, j int) bool {
	return s.by(&s.products[i], &s.products[j])
}

func main() {
	products := []Product{
		{105, "Pen", 5, 50, "Stationary"},
		{107, "Pencil", 2, 100, "Stationary"},
		{103, "Marker", 50, 20, "Utencil"},
		{102, "Stove", 5000, 5, "Utencil"},
		{101, "Kettle", 2500, 10, "Utencil"},
		{104, "Scribble Pad", 20, 20, "Stationary"},
	}

	id := func(p1, p2 *Product) bool {
		return p1.Id < p2.Id
	}
	name := func(p1, p2 *Product) bool {
		return p1.Name < p2.Name
	}
	cost := func(p1, p2 *Product) bool {
		return p1.Cost < p2.Cost
	}
	units := func(p1, p2 *Product) bool {
		return p1.Units < p2.Units
	}
	category := func(p1, p2 *Product) bool {
		return p1.Category < p2.Category
	}

	By(id).Sort(products)
	fmt.Printf("By Id:\n %v \n", products)

	By(name).Sort(products)
	fmt.Printf("By name:\n %v \n", products)

	By(cost).Sort(products)
	fmt.Printf("By cost:\n %v \n", products)

	By(units).Sort(products)
	fmt.Printf("By units:\n %v \n", products)

	//By(category).Sort(products)
	ps := &productSorter{
		products: products,
		by:       category, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
	fmt.Println("By category:", products)

}
