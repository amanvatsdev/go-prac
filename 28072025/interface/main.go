package main

import (
	"fmt"
	"math"
)

type Shape interface{
	Area()
}

func Info (x Shape){
	x.Area()
}

type square struct {
	length int
	width  int
}
func (s square) Area() {
	a:=(s.length*s.width)
	fmt.Println("the area of square is ",a)
}

type circle struct {
	radius float64
}

func (c circle) Area() {
	a:=math.Pi* math.Pow(c.radius,2)
	fmt.Println("the area of square is ",a)
}


func main() {
x:=square{
	length: 23,
	width: 23,
}

y:=circle{
	radius: 23.3,
}
x.Area()
y.Area()

fmt.Println("-----------------")

Info(x)
Info(y)
}