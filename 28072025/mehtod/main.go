package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Println("The age of",p.first," is ",p.age)
}

func main() {
a:=person{
	first: "Aman",
	age: 22,
}

a.speak()
}