package main

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func main() {

	products := []Product{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}
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
