package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	// "os"
)

type Book struct {
	Title  string
	Author string
	Price  float64
}

func main() {
	Book1 := Book{
		Title:  "Go Basics",
		Author: "Aman Vats",
		Price:  299.99,
	}

	// f, err := os.Create("text.txt")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// defer f.Close()

	var f bytes.Buffer

	fmt.Println(f)
	encoder := json.NewEncoder(&f)
	encoder.Encode(Book1)
	fmt.Println(f.String())

}
