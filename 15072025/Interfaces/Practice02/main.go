package main

import "fmt"


type animal interface{
	sound()
}

type Dog struct{}
type Cat struct{}


func (d Dog) sound() {
	fmt.Println("Bark")
}

func (c Cat) sound() {
	fmt.Println("Meow")
}

func makesound (x animal){
	x.sound()
}


func main() {
var d Dog
var c Cat

makesound(d)
makesound(c)


}