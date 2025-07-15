package main

import "fmt"

type person struct {
	FirstName    string
	LastName     string
	FavIceCreams []string
}

func main() {

	P1:=person{
		FirstName: "Aman",
		LastName: "Vats",
		FavIceCreams: []string{"Vanella","Butterscotch","Choclate"},
	}

	P2:=person{
		FirstName: "Ravi",
		LastName: "Mehra",
		FavIceCreams: []string{"Strawberry","Pista","Kesar"},
	}
	fmt.Println(P1)
	fmt.Printf("%T\t%v\n",P1,P1)

	fmt.Println("--------------------")
	
	for _, v:=range P1.FavIceCreams{
		fmt.Printf("%v\n",v)
	}
	
	m :=map[string]person{
		P1.LastName:P1,
		P2.LastName:P2,
	}
	fmt.Println(m)

	fmt.Println("-------------------------")
	
	for _,v:=range m{
		fmt.Printf("%v\n",v)
	}
}
