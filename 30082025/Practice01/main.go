package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street int    `json:"street,omitempty"`
	City   string `json:"city,omitempty"`
	State  string `json:"state,omitempty"`
}

type student struct {
	Name    string  `json:"name,omitempty"`
	Age     int     `json:"age,omitempty"`
	Address Address `json:"address,omitempty"`
}

func main() {

	jsonbyte := []byte(`[
	{"name":"Aman","age":23,"address":{"street":23,"city":"Rohtak","state":"Haryana"}},
	{"name":"Ravi","age":43,"address":{"street":22,"city":"Gurugram","state":"Haryana"}}
	]`)
	fmt.Println(string(jsonbyte))

	var jsonstring []student

	err := json.Unmarshal(jsonbyte, &jsonstring)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("stringjson value---------------")
	fmt.Println(jsonstring)
	fmt.Println("-----------------------------------------------------------")
	fmt.Printf("%+v\n", jsonstring)
	fmt.Println("-----------------------------------------------------------")

	for _, v := range jsonstring {
		fmt.Printf("%+v\n", v)
		fmt.Println("-----------------------------------------------------------")
	}

	for i, v := range jsonstring {
		fmt.Printf("Student %d:\n", i+1)
		fmt.Printf("  Name: %s\n", v.Name)
		fmt.Printf("  Age: %d\n", v.Age)
		fmt.Printf("  Street: %d\n", v.Address.Street)
		fmt.Printf("  City: %s\n", v.Address.City)
		fmt.Printf("  State: %s\n\n", v.Address.State)
	}

}
