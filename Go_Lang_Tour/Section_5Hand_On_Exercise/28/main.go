package main

import (
	"fmt"
)

/*
In this we will practice
printf verbs to show values and types
*/
func main() {
	var a, b, c = "Name", 45, 45.25
	fmt.Printf("%v of Type %T\n", a, a)
	fmt.Printf("%v of Type %T\n", b, b)
	fmt.Printf("%v of Type %T\n", c, c)
	d, e, f := "Aman", 45555, 2515.02
	fmt.Printf("%v of Type %T\n", d, d)
	fmt.Printf("%v of Type %T\n", e, e)
	fmt.Printf("%v of Type %T\n\n", f, f)
}
