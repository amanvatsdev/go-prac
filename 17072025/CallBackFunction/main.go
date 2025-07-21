package main

import "fmt"

func main() {
	a,b:=12,45
	c:=processResult(a,b,Addition)
	fmt.Println(c)
}

func processResult(a int, b int, y func(a, b int) int) int {
	return y(a, b)
}

func Addition(a, b int) int {
	return a + b
}
