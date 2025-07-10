package main

import "fmt"


func main(){
	var X [5]int
	fmt.Println(X)
	fmt.Println("lenth of x",len(X))
	X[2]=23
	fmt.Println("-------------------------")
	fmt.Println(X)
	fmt.Println("lenth of x",len(X))
}