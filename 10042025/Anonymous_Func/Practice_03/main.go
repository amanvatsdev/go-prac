package main

import "fmt"

func main() {
	A := func(x, y int) int {
		return x + y
	}
	fmt.Println(A(4,5))
}