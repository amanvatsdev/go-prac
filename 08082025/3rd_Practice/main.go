package main

import "fmt"


type youngin interface{
	walk()
	run()
}

func youngrun(x youngin){
	x.run()
	
}
type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("The name of dog is ",d.first)
}

func (d *dog) run() {
	d.first="Roger"
	fmt.Println("The name of running dog is ",d.first)
}

func main() {
x:=dog{first: "Henry"}


x.walk()
x.run()	

y:=&dog{"Charlie"}
y.walk()
y.run()
fmt.Println(y)

youngrun(&x)
youngrun(y)
}