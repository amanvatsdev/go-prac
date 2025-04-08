package main

import "fmt"

func IdentifyType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("this is of string type ",v)
	case int:
		fmt.Println("this is of int type ",v)
	case float64:
		fmt.Println("this is of float64 type ",v)
	case bool:
		fmt.Println("this is of bool type ",v)
	default :
		fmt.Println("Unknown Type")
	}
}
func main() {
	var A interface{}

	A =45
	IdentifyType(A)
	A="Aman"
	IdentifyType(A)
	A=true
	IdentifyType(A)
	A=33.3
	IdentifyType(A)
	
	
}