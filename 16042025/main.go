package main

import "fmt"

func PerformTask(x,y int,callback func(int,int)int){
	fmt.Println(callback(x,y))
}
func main() {
	Addition:=func (x ,y int)int  {
	return x+y
	}
	Subraction:=func(x,y int)int{
		return x-y
	}
	Multiplication:=func(x,y int)int{
		return x*y
	}
	Division:=func(x,y int)int{
		return x/y
	}
	fmt.Println("plz give me two number one by one")
	var N1 int
	fmt.Scanln(&N1)
	var N2 int
	fmt.Scanln(&N2)
	fmt.Println("Now plz tell me the task you want to perform plz select an option A,B,C,D")
	fmt.Println(`
	A=Additon,
	B=Subtraction
	C=Multiplication
	D=Division`)
	var task string
	fmt.Scanln(&task)
	
	switch  {
	case task=="A"||task=="a":
		PerformTask(N1,N2,Addition)
	case task=="B"||task=="b":
		PerformTask(N1,N2,Subraction)
	case task=="C"||task=="c":
		PerformTask(N1,N2,Multiplication)
	case task=="D"||task=="d":
		PerformTask(N1,N2,Division)
	default:
		fmt.Println("Choose the valid option")
	}
	//R53282682== user id
	//5085461187312773== debit card

	// a, b := 1, 2
	// fmt.Println(a,b)
	// fmt.Println(Addition(2,2))
	// PerformTask(34,34,Addition)

}