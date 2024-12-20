/*
write a program that declare two values

	1.one variable to store a value of type int8
		-assign to it maximum Number Possible ,then print it
	2.one variable to store a value of type uint8
		-assign to it maximum Number Possible ,then print it
*/
package main

import "fmt"

func main() {
	a := 128
	var b uint = 256
	fmt.Printf("a of type :%T \twith value:%d\n", a, a)
	fmt.Printf("b of type :%T\t with value:%d\n", b, b)
}
