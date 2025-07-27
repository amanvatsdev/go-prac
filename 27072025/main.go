package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	x := 0
	return x
}

func bar() (int, string) {
	x := 0
	name := "Aman"
	return x, name
}