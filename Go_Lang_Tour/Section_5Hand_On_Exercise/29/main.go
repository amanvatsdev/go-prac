/*
in this
Write a program that uses printd verb	and show the following numberes
1=747
2=911
3=90210
as

	Decimal
	Binary
	Hexamdecimal
*/
package main

import (
	"fmt"
)

func main() {
	a, b, c := 747, 911, 90210
	fmt.Printf("Decimal value:%d,Binary Value:%b,hexadecimal:%x", a, a, a)
	fmt.Printf("Decimal value:%d,Binary Value:%b,hexadecimal:%x", b, b, b)
	fmt.Printf("Decimal value:%d,Binary Value:%b,hexadecimal:%x", c, c, c)

}
