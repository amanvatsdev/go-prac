package main

import "fmt"

func main() {
	a:=func() {
		for x := 0; x <= 10; x++ {
			fmt.Println(x)
		}
	}
	a()

	fmt.Println("------------------------------------")
	b:=add(6,7)
	fmt.Println(b)
	fmt.Println("------------------------------------")
	
	c:=quick()
	fmt.Println("c:",c())
}

func add (a,b int)int{
	return a+b
}

func quick()(func()int){
	 return func() int {
		return 42	
	 }
}