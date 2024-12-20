package main

import "fmt"

func main() {
	//In this Use of Binary and Hexadecimal in coding

	A := 45
	B := "Aman Vats"

	fmt.Printf("Binar Value :=%b\nHexadecimal value :=%x\n", A, A)
	fmt.Println("--------------------------------------------")
	fmt.Println("Binary :")
	for _,c:=range B{
		fmt.Printf("%08b ",c)
	}
}
