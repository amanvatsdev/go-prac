package main

import "fmt"

func main() {
	func(a int) {
		Y := a * a
		fmt.Println("Sqare of ", a, "is",Y)
	}(5)
}
