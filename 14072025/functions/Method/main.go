package main

import "fmt"


type person struct{
	name string
	Age int
}
func main() {
	//variadic parameter

	x := []int{23, 32, 345, 345, 456, 3, 32, 23, 3}
	fmt.Println(Sum(x...))		//in this when we use ...after variable x we UNFURL A SLICE 
	fmt.Println(x)

	
	fmt.Println("----------------------------Mehtod-------------------------------")
	p1:=person{
		name: "Raju",
		Age:24,
	}

	P2:=person{
		name: "Rani",
		Age:23,
	}

	p1.Speak()
	P2.Speak()
	
}

func Sum(x ...int) int {
	n := 0
	for _, b := range x {
		n+=b
	}
	return n

}

func (p person)Speak(){
	fmt.Println("My name is ",p.name," and my age is ",p.Age)
}