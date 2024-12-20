package main

import (
	import2 "aman/Import_02"
	Import "aman/import_01"
	"fmt"
)

var y string = "fjfjfjfjfj" //this y is on package level so it can be used at any level also in  other functions

func Random() {
	x := 45
	fmt.Println(x)
	//Here the scope is x is only limited to Random Function
}
func main() {
	Import.DistanceBetweenSunAndEarth()
	import2.DistanceBetweenEarthandMoon()
	fmt.Println("---------------------------")
	fmt.Println(Import.SunToMoon)
	//In this code we learned that how scope works in go lang
	fmt.Println(y) 
	fmt.Println(y)
}
