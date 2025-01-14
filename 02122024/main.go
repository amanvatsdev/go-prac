package main

import "fmt"

func main() {
	var N int
	fmt.Println("Give me a number of which you want a table:")
	fmt.Scanln(&N)

	for multiplier := 1; multiplier <= 10; multiplier++ {
		fmt.Println(N, "x", multiplier, "=", N*multiplier)
		fmt.Println("---------------------")
	}

}