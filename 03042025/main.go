package main

import "fmt"

type data struct {
	name string
}


type SpecialData struct{
	data
	age int
}
func (d data) Info() {
	fmt.Println("My name is ",d.name)
}

func (s SpecialData)Info(){
	fmt.Println("My Name is ",s.name ,"and my age is ",s.age)
}



func main() {
	D1:=data{
		name: "Aman",
	}

	D2:=SpecialData{
		data:data{
			name: "Raj",
		},
		age: 22,
		}
	D1.Info()
	D2.Info()
	D2.Info()
}