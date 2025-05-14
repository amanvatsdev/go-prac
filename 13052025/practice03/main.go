package main

import "fmt"

func makeMultiplier(Number int) func(int) int {
	Factor:=0
	return func(i int) int {
		Factor=Number*i
		return Factor
	}
}

func main() {
	value1:=makeMultiplier(5)
	fmt.Println(value1(1))
	fmt.Println(value1(2))
	fmt.Println(value1(3))
}