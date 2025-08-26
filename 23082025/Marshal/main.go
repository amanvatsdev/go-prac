package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func main() {
	persons:=[]person{
		{Name: "Aman",Age: 22,},
		{Name: "Rahul"},
	}
	

	jsonData,err:= json.Marshal(persons)
	if err != nil{
		fmt.Println("error found :",err)
	}
	fmt.Println(string(jsonData))

}

