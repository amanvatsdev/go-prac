package main

import "fmt"

// func Product(x int ,func(int)) {
// 	func(y int) {
// 	return	fmt.Println("Product is ",x+y)
// 	}
// }

func Product(x int)func (int){
	return func(i int) {
		z:=x*i
		fmt.Println(z)
	}
}

func main() {
	sum:=Product(34)
	sum(23)
	
}