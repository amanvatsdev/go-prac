package main

import "fmt"

func main() {
	x := []string{"Aman", "Raju", "Ravi", "Pawan", "Bhawna", "Satish"}
	fmt.Println(x)
	fmt.Println("Lenth",len(x))
	fmt.Println("-----------------------")
	
	for i,v:=range x{
		fmt.Println(i,v)
	}
	fmt.Println("=====================")
	//slicing a slice
	
	fmt.Printf("%v\n",x[1:5])
	
	fmt.Printf("%v\n",x[2:])
	fmt.Printf("%v\n",x[:4])
	fmt.Printf("%v\n",x[:])
	
	//deleting a slice

	x=append(x[:2] , x[3:]...)
	
	fmt.Printf("%v\n",x[:])
	fmt.Println("=====================")
	
	//mutlidimentional slice
	
	y:=[]string{"Aman","Vats"}
	z:=[]string{"Naresh","Kumar"}

	fmt.Println(y)
	fmt.Println(z)
	
	A:=[][]string{x,y,z}
	fmt.Println("A:",A)
	fmt.Println("-------------------------------")
}