package main

import (
	"fmt"
)

func main() {
	Name := []string{"Aman", "Sinu", "Vikash", "Harsh"}
	
	for index, value := range Name {
		fmt.Printf("index=%d\tvalue=%s    =",index,value)
		if value=="Sinu"{
			fmt.Println("He is my Brother",)
		}else{
			fmt.Println("I dont know him.")
		}
		
		}
	}

