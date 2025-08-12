package main

import (
	"fmt"
	"math"
)

type shape interface {
	area()
	perimeter()
}

type circle struct {
	radius float64
}

func (c circle) area() {
	x := math.Pi*c.radius * c.radius
	fmt.Println("The area of circle is",x)
}
func (c circle)perimeter(){
	x:=2*math.Pi*c.radius
	fmt.Println("The perimeter of circle is", x)
}

type rectangle struct {
	length float64
	breath float64
}

func (r rectangle)area(){
	x:=r.length*r.breath
	fmt.Println("The area of rectangle is",x)
}
func (r rectangle)perimeter(){
	x:=2*(r.length+r.breath)
	fmt.Println("The perimeter of rectangle is",x)
}

func Printer(x shape){
	x.area()	
	x.perimeter()
}
func main() {
	r1:=rectangle{
		length: 2,
		breath: 3,
	}
	c1:=circle{
		radius: 23.5,
	}
	// r1.area()
	// c1.area()
	Printer(r1)
	Printer(c1)

}
