package main

import "fmt"

func main() {
	x:=counter()
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}