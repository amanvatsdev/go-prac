package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func addf(a, b float64) float64 {
	return a + b
}

func addT[T int | float64](x, y T) T {
	return x + y
}

func main() {
	a, b := 4, 6
	c, d := 2.4, 3.7

	x:=add(a, b)
	y:=addf(c, d)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(addT(a,b))
	fmt.Println(addT(c,d))
}