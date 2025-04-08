package main

import "fmt"

type Geometry interface {
	Area() float64
	Perimeter() float64
}

type circle struct {
	radius float64
}

func (c circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}
func (c circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

type triangle struct {
	side1Base float64
	side2     float64
	side3     float64
	height    float64
}

func (t triangle) Area() float64 {
	return 0.5 * t.side1Base * t.height
}

func (t triangle) Perimeter() float64 {
	return t.side1Base + t.side2 + t.side3
}

// func PrintGeometry(g Geometry){
// 	fmt.Println("Area",g.Area())
// 	fmt.Println("Perimeter",g.Perimeter())
// }
func main() {
	var g Geometry

	C1 := circle{
		radius: 23,
	}

	g = C1

	fmt.Println(g.Area())
	fmt.Println(g.Perimeter())
	fmt.Println("---------------------------------------")
	fmt.Println(g)
	fmt.Println("---------------------------------------")
	T1 := triangle{
		side1Base: 5,
		side2:     6,
		side3:     4,
		height:    3.97,
	}

	fmt.Println(T1.Area())
	fmt.Println(T1.Perimeter())

	fmt.Println(C1.Area())
	fmt.Println(C1.Perimeter())
	// PrintGeometry(C1)
	// PrintGeometry(T1)

}
