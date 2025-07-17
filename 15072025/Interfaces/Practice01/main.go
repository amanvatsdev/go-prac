package main

import "fmt"

type person struct{
	name string
}

func (p person)speak(){
	fmt.Println("my name is ",p.name)
}

type secretAgent struct{
	person
	ltk bool
}

func (s secretAgent)speak(){
	fmt.Println("I am a secret Agent and my name is ",s.name)
}

type Unknown interface{
	speak()
}

func saySomething(x Unknown){
	x.speak()
}


func main() {

p1:=person{
	name: "Aman",
}
Sa1:=secretAgent{
	person: person{
		name: "Raju",
	},
	ltk: true,
}

p1.speak()
Sa1.speak()

fmt.Println("----------------------")

saySomething(p1)
saySomething(Sa1)

}