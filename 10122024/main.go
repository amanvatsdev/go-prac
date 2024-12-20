package main

import (
	"fmt"
	"math"
)

type Formula struct {
	Radius float64
	Height int
}

func (a Formula) Area() float64 {
	return math.Pi * a.Radius * a.Radius
}
func main() {
	Circle := Formula{
		Radius: 6,
	}
	fmt.Println(Circle.Area())
}
