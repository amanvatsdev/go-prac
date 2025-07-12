package main

import (
	"fmt"
)

func main() {
	x := [5]string{}
	x[0] = "aman"
	x[2] = "Ravi"
	x[3] = "Raju"
	x[4] = "Renu"
	x[1] = "Geeta"
	fmt.Println(x)

	fmt.Println("--------------------------------")

	for i, v := range x {
		fmt.Printf("%v\t%T\t%v\n", i, v, v)
	}
	fmt.Println("--------------------------------")

	y := []int{}

	y = append(y, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51)

	for i, v := range y {
		fmt.Printf("%v\t%T\t%v\n", i, v, v)
	}
	fmt.Printf("%T\t%v\n", y, y)
	fmt.Printf("%T\t%#v", y, y)
	//slicing a slice
	fmt.Println("--------------------------")
	fmt.Println(y[:5])
	fmt.Println(y[5:])
	fmt.Println(y[2:7])
	fmt.Println("--------------------------")

	z := []int{42, 23, 44, 45, 46, 47, 48, 49, 50, 51}

	z = append(z, 52)

	fmt.Println(z)

	z = append(z, 53, 54, 55)

	fmt.Println(z)
	z = append(z, y...)
	fmt.Println(z)

	z = append(z[:14], z[24:]...)
	fmt.Println(z)

}
