package main

import "fmt"

func main() {

	a, b := 23, 45

	z:=getMultiplier(a)
	fmt.Println(z(b))
}

func getMultiplier(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}