package main

import (
	"fmt"
	"math/rand"
)

func main() {
	random := rand.Intn(250)
	fmt.Println("rand :=", random)

	if random <= 100 {
		fmt.Println("No is bwtween 0 to 100")
	} else if random > 100 && random <= 200 {
		fmt.Println("No is between 101 to 200")
	} else if random > 200 && random <= 250 {
		fmt.Println("No is between 201 to 250")
	} else {
		fmt.Println("Error while printing the number")
	}

	fmt.Println("-------------------")
	for {
		x := rand.Intn(4)
		fmt.Println(x)
		if x == 3 {
			break
		}
	}

	fmt.Println("-------------------------")

	for {
		x := rand.Intn(8-3+1) + 3
		fmt.Println(x)
		if x == 8 {
			break
		}
	}
	fmt.Println("===============================")
	y := 25
	for y >= 0 {
		fmt.Println(y)
		y--
	}
	z := 1
	for z <= 25 {
		fmt.Println(z)
		z++
	}
	fmt.Println("----------------------------------------")
	for x := 0; x <= 30; x++ {
		if x%2 != 0 {
			fmt.Println(x)
			} else {
			continue
		}
	}
	fmt.Println("----------------------------------------")

	for x:=1;x<=5;x++{
		for d:=1;d<=5;d++{
			fmt.Printf("Outer loop %v\t innner loop %v \n",x,d)
		}
	}
	fmt.Println("--------------------------------------")
	
	x:=[]int{34,23,43,34,6,34,45}
	for i,v:=range x{
		fmt.Printf("%d:%d\n",i,v)
	}
	fmt.Println("--------------------------------------")
	
	m:=make(map[string]int)
	
	addallmap(m,map[string]int{"aman":1,"ravi":2,"Radhe":3})

	for name,int:=range m{
		fmt.Printf("%s:%d\n",name,int)
	}
}


func addallmap(m map[string]int, v map[string]int){
	for s,i:= range v{
		m[s]=i
	}
}