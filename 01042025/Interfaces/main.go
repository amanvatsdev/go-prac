package main

import "fmt"


type rectangle struct{
	lenght int
	width  int
}

func (r rectangle)Area()int{
	return r.lenght * r.width
	
}

func main(){
	R1:=rectangle{
		lenght: 5,
		width: 4,
	}
	fmt.Print("Area of Rectangle ",R1.Area())
}






















// type data struct {
// 	name string
// 	age  int
// }

// func (d data) Detail()string {
// 	return fmt.Sprintf("my name is %s and my age is %d",d.name,d.age)
// }

// func main(){
// 	a:=data{
// 		name: "aman",
// 		age: 24,
// 	}

// fmt.Println(a.Detail())

// }