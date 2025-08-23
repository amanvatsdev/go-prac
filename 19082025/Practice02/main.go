package main

import "fmt"

type mynumbers interface {
	int | float64
}

func addT[T mynumbers](x, y T) T {
	return x + y
}

func main() {
	a, b := 4, 6
	c, d := 2.4, 3.7

	fmt.Println(addT(a, b))
	fmt.Println(addT(c, d))
}
