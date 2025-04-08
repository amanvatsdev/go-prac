package main

import "fmt"

type Shape interface {
	Area()float64
}

type rectangle struct {
	Width  int
	Breath int
}

type circle struct {
	radius int
}

func (r rectangle) Area()float64 {
	s := r.Breath * r.Width
	return float64(s)
	// fmt.Println("Area of Rectangle is ", s)
}

func (c circle)Area()float64{
	s:=3.14 * float64(c.radius) *float64(c.radius)
	return s
	// fmt.Println("Area of Circle is ",s)
}

func AreaCalculator(s Shape){
	fmt.Println(s.Area())
}

func main() {
	S1 := rectangle{
		Breath: 24,
		Width:  25,
	}
	S2:=circle{
		radius: 23,
	}
	S1.Area()
	S2.Area()
	// AreaCalculator(S1)
	// AreaCalculator(S2)
	AreaCalculator(S1)

}
