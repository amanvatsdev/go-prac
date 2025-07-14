package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make   string
	model  string
	doors  int
	colour string
}

func main() {
	V1 := vehicle{
		engine: engine{true},
		make:   "Mahindra",
		model:  "Thar",
		doors:  5,
		colour: "Black",
	}

	V2 := vehicle{
		engine: engine{electric: false},
		make:   "Hyuandey",
		model:  "Verna",
		doors:  4,
		colour: "White",
	}

	fmt.Println(V1)
	fmt.Println(V2)

	fmt.Println("-------------------")

	fmt.Println(V1.model  ,V1.engine.electric)
	fmt.Println(V2.model  ,V2.engine.electric)

a1:=struct{
	first string
	friends map[string]int
	favDrinks []string
}{
	first: "Aman",
	friends: map[string]int{
		"Aarush":23,
		"Nagender":26,
		"Dinesh":32,
	},
	favDrinks: []string{"Coca","Dew","Sprite"},
}
fmt.Println("----------------------------------------------")
fmt.Println(a1)

for K,t:=range a1.favDrinks{
	fmt.Printf("%v\t%v\n",K,t)
}

}
