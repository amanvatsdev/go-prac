package main

import (
	"fmt"
	"log"
	
)

func Handleinput(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("this is of string type ", v)
	case int:
		fmt.Println("this is of int type ", v)
	case float64:
		fmt.Println("this is of float64 type ", v)
	case bool:
		log.Println("this is of bool type ", v)
	default:
		fmt.Println("Unknown Type")
	}
}


func main() {
	var A interface{}

	A =45
	Handleinput(A)
	A="Aman"
	Handleinput(A)
	A=true
	Handleinput(A)
	A=33.3
	Handleinput(A)


fmt.Println("---------------------------------------------------")
B:=[]interface{}{"Aman",36,true,23.443}
	
	for _,values:=range B{
		switch v:=values.(type){
		case string:
			fmt.Println("this is of string type ", v)
		case int:
			fmt.Println("this is of int type ", v)
		case float64:
			fmt.Println("this is of float64 type ", v)
		case bool:
			fmt.Println("this is of bool type ", v)
		default:
			fmt.Println("Unknown Type")
		}
	}
}
