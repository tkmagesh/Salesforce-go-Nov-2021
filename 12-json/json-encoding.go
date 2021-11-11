package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	Id       int      `json:"id" xml:"prod-id"`
	Name     string   `json:"name"`
	Cost     float32  `json:"cost"`
	Units    int      `json:"units"`
	Category Category `json:"category"`
}

func main() {
	stationaryCategory := Category{Id: 101, Name: "Stationary"}
	utencilCategory := Category{Id: 102, Name: "Utensils"}
	electronicsCategory := Category{Id: 103, Name: "Electronics"}

	products := []Product{
		{105, "Pen", 5, 50, stationaryCategory},
		{103, "Pencil", 2, 100, stationaryCategory},
		{102, "Marker", 50, 20, stationaryCategory},
		{104, "Frying Pan", 500, 5, utencilCategory},
		{101, "Phone", 5000, 3, electronicsCategory},
		{100, "Bowl", 100, 50, utencilCategory},
	}

	// Encode the products as JSON
	file, err := os.Create("products.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if e := encoder.Encode(products); e != nil {
		log.Fatalln(e)
	}
	fmt.Println("Done")
}
