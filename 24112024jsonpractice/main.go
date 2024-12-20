package main

import (
	"encoding/json"
	"fmt"
)

type Details struct {
	Name    string `json:"Naam"`
	Age     int    `json:"Umar"`
	Address string `json:"Pata"`
}

func main() {

	Person1 := Details{
		Name:    "Sinu",
		Age:     25,
		Address: "Chiri",
	}
	Person2 := Details{
		Name:    "Aman",
		Age:     21,
		Address: "Badli",
	}
	Persons_Data := []Details{Person1, Person2}
	fmt.Println("This is the simple Form Of Data", Persons_Data)
	//Coverting go lang data into json data type
	vikash, err := json.Marshal(Persons_Data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("This the converted data from go lang type to json :", string(vikash))

	//Now conver jsn Data into go lang data type

	var unmarshal []Details
	err = json.Unmarshal(vikash, &unmarshal)
	fmt.Println("this is again main data :", unmarshal)
}
