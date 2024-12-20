package main

import (
	"fmt"
)

/*
In This We learned use of following concepts
zero value, :=, type specificity, blank identity
*/

func main() {
	var a int //Zero Value
	fmt.Println(a)
	var b string = "Wealth"
	fmt.Println(b)
	c := 45 //Short VAriable Declaration(:=)
	d := "Important"
	fmt.Println(c, d)
	e, f, g, _ := 48, 45.23, "Hi", 48 //blank value
	fmt.Printf("%v of type %T\n", e, e)
	fmt.Printf("%v of type %T\n ", f, f)
	fmt.Printf("%v of type %T\n ", g, g)

}
