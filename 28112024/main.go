package main

import "fmt"

func Average(numberes ...int) int {
	sum := 0
	for _, total := range numberes {
		sum += total
	}

	avg := sum / len(numberes)
	if avg == 0 {
		panic("Average is zero")
	}
	return avg

}
func main() {
	fmt.Println(Average(0, 1, 2, 2, 1, 5, 5, 5))
}

// package main

// import "fmt"

// func main() {
// 	Addition := func(x, y int) int {
// 		return x + y
// 	}
// 	fmt.Println(Addition(99, 101))
// }

// package main
// import(
// 	"fmt"
// )
// func Add (Aman ...int)int{
// 	Total:=0
// 	for _,Sum :=range Aman{
// 		Total += Sum
// 	}
// 	return Total

// }

// func main(){
// 	fmt.Println(Add(5,6,4,5,6,4,6))

// }
