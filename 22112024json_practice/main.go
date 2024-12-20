package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"Naam"`
	Age  int    `json:"umar"`
	Sex  string `json:"ling"`
}

func main() {
	People := person{
		Name: "Aman",
		Age:  25,
		Sex:  "Male",
	}

	// converting data into json
	fmt.Println("converting data into json")
	vikash, sinu := json.Marshal(People)
	if sinu != nil {
		fmt.Println("This the error :", sinu)
	}

	fmt.Println(string(vikash))

	fmt.Println("converting json data back into a variable")

	var a person

	json.Unmarshal(vikash, &a)
	fmt.Println(a)

}
