// This is package main
package main

import "fmt"

//Main funciton
func main() {
	x, y := "Aman", "Raju"
	fmt.Println(Intro(x, y))
}

//INro will printout the name of two students.
func Intro(x, y string) string {
	return fmt.Sprintf("The name of students are %s and %s", x, y)
}
