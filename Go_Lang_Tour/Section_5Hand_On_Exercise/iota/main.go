package main

import "fmt"

func main() {
	const (
		a = 15 + iota
		b
		c
		d
		_ //this skips the constant iota
		e
	)
	fmt.Printf("%d\t%T\n", a, a)
	fmt.Printf("%d\t%T\n", b, b)
	fmt.Printf("%d\t%T\n", c, c)
	fmt.Printf("%d\t%T\n", d, d)
	fmt.Printf("%d\t%T\n", e, e)
}
