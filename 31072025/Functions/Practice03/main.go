package main

import "fmt"

func main() {
	a := Incrementor(0)
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}

func Incrementor(y float64) func() float64 {
	var x float64
	return func() float64 {
		x++
		return x
	}
}
