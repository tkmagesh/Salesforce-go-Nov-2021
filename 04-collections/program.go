package main

import "fmt"

func main() {
	//Array
	fmt.Println("Array")

	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	//var nos [5]int = [...]int{3, 1, 4, 2, 5}
	var nos = [5]int{3, 1, 4, 2, 5}

	fmt.Println(nos)

	fmt.Println("Iterating using an indexer")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Println(nos[idx])
	}

	fmt.Println("Iterating using range")
	/*
		for idx, value := range nos {
			fmt.Println(value, idx)
		}
	*/

	for _, value := range nos {
		fmt.Println(value)
	}

	fmt.Println("Copying an array")
	var newNos = nos
	newNos[0] = 100
	fmt.Printf("nos = %v, newNos = %v\n", nos, newNos)

	//Slice
	fmt.Println("Slice")
	var products []string
	//var products []string = []string{"Laptop", "Mobile", "TV", "Camera"}
	fmt.Println(products)

	/*
		products = append(products, "Laptop")
		products = append(products, "Mobile")
	*/

	products = append(products, "Laptop", "Mobile")
	fmt.Println(products)

	stationaryProducts := make([]string, 2)
	stationaryProducts[0] = "Pen"
	stationaryProducts[1] = "Pencil"
	//stationaryProducts := []string{"Pen", "Pencil"}
	products = append(products, stationaryProducts...)
	fmt.Println(products)

	fmt.Println("Len & capacity")
	fmt.Printf("len(stationaryProducts) = %d, cap(stationaryProducts) = %d\n", len(stationaryProducts), cap(stationaryProducts))
	fmt.Println("After appending 'Marker' to stationary products")
	stationaryProducts = append(stationaryProducts, "Marker")
	fmt.Printf("len(stationaryProducts) = %d, cap(stationaryProducts) = %d\n", len(stationaryProducts), cap(stationaryProducts))
	stationaryProducts = append(stationaryProducts, "Scribble-Pad")
	fmt.Printf("len(stationaryProducts) = %d, cap(stationaryProducts) = %d\n", len(stationaryProducts), cap(stationaryProducts))
	stationaryProducts = append(stationaryProducts, "Ruler")
	fmt.Printf("len(stationaryProducts) = %d, cap(stationaryProducts) = %d\n", len(stationaryProducts), cap(stationaryProducts))

	/* slicing */
	/*
		[lo : hi] => from lo to hi-1
		[lo : ] => from lo to end
		[ : hi] => from 0 to hi-1
		[lo : lo] => empty slice
		[lo : lo + 1] === [lo] => data at lo
		[:] => copy of the slice
		[lo : hi step step] => from lo to hi-1 with step
	*/
	fmt.Println("Slicing")
	fmt.Println("Stationary Products => ", stationaryProducts)
	fmt.Println("Stationary Products[1:3] => ", stationaryProducts[1:3])
	fmt.Println("Stationary Products[2 :] => ", stationaryProducts[2:])
	fmt.Println("Stationary Products[:2] => ", stationaryProducts[:2])

	fmt.Println("Maps")
	var productRankings = map[string]int{
		"Pen":    4,
		"Pencil": 2,
		"Marker": 1,
	}

	fmt.Println(productRankings)
	fmt.Println("Ranking of Pen => ", productRankings["Pen"])

	fmt.Println("Appending a new key value pair")
	productRankings["Scribble-Pad"] = 3
	fmt.Println(productRankings)

	fmt.Println("Checking if a key exists")
	if rank, exists := productRankings["Ruler"]; !exists {
		fmt.Println("Ruler does not exist")
	} else {
		fmt.Println("Ruler exists => ", rank)
	}

	fmt.Println("Deleting a key value pair")
	delete(productRankings, "Pen")
	fmt.Println(productRankings)

	fmt.Println("Iterating over a map")
	for key, value := range productRankings {
		fmt.Println(key, value)
	}
}
