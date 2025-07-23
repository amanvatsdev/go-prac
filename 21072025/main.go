package main

import "fmt"

func main() {

	evenPrinter(4)
}

func evenPrinter(n int) {
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
