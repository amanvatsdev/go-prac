package main

import (
	"fmt"
	practice "merger/13052025/practice01"
	"merger/13052025/practice02"
)

func add(x, y int) int {
	z := x + y
	return z
}

func greetuser(name string) string {
	return fmt.Sprintln("Hello", name)
}

func fiftyoff(x int) int {
	return x - 50
}

func applydiscount(amount int, discount func(int) int) {
	a := discount(amount)
	fmt.Println("This is amount after applying disount ", a)
}

func main() {
	Sum := add(45, 23)
	fmt.Println(Sum)
	greet := greetuser("Aman")
	fmt.Println(greet)
	applydiscount(230, fiftyoff)
	area := practice.Calculate(5, practice.Square)
	fmt.Println("Calculated area is:", area)

	discount := practice02.DiscountCalculator()
	fmt.Println(discount(1005))

	// Annonymous fucntion
	multipleNo:=func(x,y int)int{
		return x*y
	}
	fmt.Println(multipleNo(5,7))


	//Immediately Invoked Function Expression

	func(name string){
		fmt.Println("Hello",name)
	}("Aman")

}
