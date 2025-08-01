package main

import "fmt"

func main() {
	x := 2
	fmt.Println("x",x)
	fmt.Println("x",x)
	p:=&x
	fmt.Println("p",p)
	fmt.Println(*p)
	fmt.Println("-------------------------")
	*p=3
	fmt.Println("p",p)
	fmt.Println(*p)


	num:=42
	numPtr:=&num
	fmt.Println("-----------------------------------")
	fmt.Printf("Type:%T and value:%v \n",num,num)
	fmt.Printf("Type:%T and value:%v ",numPtr,numPtr)
}
