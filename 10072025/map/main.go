package main

import "fmt"

func main() {

	x := map[int]string{}

	x[1] = "Aman"
	x[2] = "Arun"
	x[3] = "Annu"

	fmt.Println("len",len(x))
	fmt.Println(x)
	y:=map[string]int{
		"Aman":23,
		"Raju":24,
		"Durgesh":34,
	}
	fmt.Println(y)
	fmt.Println("len",len(y))
	
	z:=make(map[int]string)
	
	z[2]="Dubbe"
	z[5]="dinesh"
	z[3]="Dhawan"
	fmt.Println(z)
	fmt.Println("lenth",len(z))
	
	for i,v:= range z{
		fmt.Printf("%v\t%T\t%v\n",i,v,v)
	}
	fmt.Println("----------------------")
	fmt.Println(z)
	delete(z,2)
	fmt.Println(z)
	fmt.Println("----------------------")


	a:=map[string][]string{
		"Aman":{"Shaken,not stirred","martinis","fast cars"},
		"MoneyPenny_Jenny":{"intelligence","literature","computer Science"},
		"No_dr":{"Cats","Icecream","Sunset"},
	}

	for i,v:=range a{
		fmt.Printf("Name:%v\n",i)
		for index,value:=range v{
			index ++
			fmt.Printf("Hobbies:%v.%v\n",index,value)
		}
	}

}