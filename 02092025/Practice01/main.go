package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Create("Aman.txt") // f = loction of f
	defer f.Close()

	value, _ := f.Write([]byte("My Name is Aman\n")) //value = 15
	fmt.Println("f ", f)
	fmt.Println("value:", value)
	fmt.Println(f.Write([]byte("Hello world from Aman\n")))
	fmt.Fprintln(f, "Greetings to all")
}
