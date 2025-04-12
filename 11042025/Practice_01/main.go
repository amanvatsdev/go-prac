	package main

	import "fmt"

	func main() {
		X := func(x int) func(int) int {
			return func(y int) int {
				return x * y
			}
		}
		Random := X(4)(3)
		fmt.Println(Random)
	}