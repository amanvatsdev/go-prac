package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string  `json:"title,omitempty"`
	Author string  `json:"author,omitempty"`
	Price  float64 `json:"price,omitempty"`
}

func main() {
	books := []Book{
		{
			Title:  "The Alchemist",
			Author: "Paulo Coelho",
			Price:  230.3,
		},
		{
			Title:  "1984",
			Author: "George Orwell",
			Price:  823.52,
		},
		{
			Title:  "The Hunger Games",
			Author: "Suzanne Collins",
			Price:  983,
		},
	}

	for i, v := range books {
		
		fmt.Println("Book no.",i+1)
		fmt.Printf("Title:%s\nAuthor:%s\nPrice:%.1f\n\n",v.Title,v.Author,v.Price)
	}


	jsonData1,err:=json.MarshalIndent(books,"","   ")
	if err != nil {
		fmt.Println("error ",err)
		return
	}
	fmt.Println("Json converted data------------------------------")
	fmt.Println(string(jsonData1))

	 var jonBook []Book

	 errr:=json.Unmarshal(jsonData1,&jonBook)
	 if errr!=nil{
		fmt.Println(errr)
		return
	 }
	 fmt.Println("-----------------------------------------------")
	//  fmt.Println(jonBook)
	 for i, v := range jonBook {
		
		fmt.Println("Book no.",i+1)
		fmt.Printf("Title:%s\nAuthor:%s\nPrice:%.1f\n\n",v.Title,v.Author,v.Price)
	}

}
