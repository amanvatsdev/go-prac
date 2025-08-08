package main

import "fmt"

func givePointer() *int {
	x := 23
	return &x
}

func addStackValues()int{
	a:=23
	b:=42
	return a+b
}

func main() {
	output := givePointer()
	fmt.Println(*output)
	fmt.Println("========================")
	fmt.Println(addStackValues())
}