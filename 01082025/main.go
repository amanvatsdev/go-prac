package main

import "fmt"

func isDelta(a *int) {
	*a = 53
}

func sliceDelta(x []int) {
	x[1] = 23
}

func mapDelta(x map[string]int,y string){
	x[y]=23
}

func main() {
	a := 23
	fmt.Println(a)
	b := &a
	fmt.Println("b", b)
	isDelta(&a)
	c := &a
	fmt.Println("a", a)
	fmt.Println("c", *c)

	fmt.Println("---------------------------")
	x := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(x)
	sliceDelta(x)
	fmt.Println(x)

	fmt.Println("---------------------------")
	z := map[string]int{
		"Aman":  22,
		"Rahul": 23,
	}
	fmt.Println(z)
	mapDelta(z,"Aman")
	fmt.Println(z)
	fmt.Println("---------------------------")

}
