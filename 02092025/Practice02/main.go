package main

import (
	"fmt"
	"os"
)

func main() {
	x, _ := os.Open("text.txt")
	buf := make([]byte, 50)
	m, _ := x.Read(buf)
	fmt.Println(m)
	fmt.Println(string(buf[:m]))
}
