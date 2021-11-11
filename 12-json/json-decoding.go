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

	products := []Product{}

	// Encode the products as JSON
	file, err := os.Open("products.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	encoder := json.NewDecoder(file)
	if e := encoder.Decode(&products); e != nil {
		log.Fatalln(e)
	}
	fmt.Println(products)
	fmt.Println("Done")
}
