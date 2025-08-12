package main

import "fmt"

func main() {
	x, y := 23, "Rahul"

	a, b := &x, &y

	fmt.Println("a",a)
	fmt.Println("b",b)
	fmt.Println("x",x)
	fmt.Println("y",y)
	fmt.Println("0--------------------------------------------0")
	fmt.Printf("type of x:%T, y:%T\n",x,y)
	fmt.Printf("type of a:%T, b:%T\n",a,b)
	fmt.Println("0--------------------------------------------0")
	fmt.Println("a",*a, "b",*b)
}