package main

import (
	"fmt"
)

func main() {
	a:=9
	b:=PrintSquare(square,a)
	fmt.Println(b)
}

func PrintSquare(a func(int)int ,b int)string{
	c:=a(b)
	return fmt.Sprintf("the number is %d and the square is %d",b,c)
}
func square(x int) int {
	return x*x
}