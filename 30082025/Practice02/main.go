package main

import (
	"encoding/json"
	"fmt"
)

type todo struct {
	Title   string `json:"title,omitempty"`
	Done    bool   `json:"done,omitempty"`
	DueDate string `json:"duedate,omitempty"`
}

func main() {
	jsondata := []byte(`[
	{"title":"Learn Go","done":false,"duedate":"2025-08-25"},
  	{"title":"Go to Gym","done":true,"duedate":"2025-08-26"},
  	{"title":"Finish SQL Practice","done":false,"duedate":"2025-08-30"}
	]`)

	var jsonstring []todo

	err := json.Unmarshal(jsondata,&jsonstring)
	if err!=nil{
		fmt.Println("error:",err)
	}

	for i,v:= range jsonstring{
		if !v.Done{
			fmt.Println("Student:",i+1)
		fmt.Printf("Title :%+v\n",v.Title)
		fmt.Printf("Done:%+v\n",v.Done)
		fmt.Printf("DueDate:%+v\n\n",v.DueDate)
		}
	}


	

}
