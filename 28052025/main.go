package main

import "fmt"

func CliApp(N1, N2 float64, operation string) {

	switch operation {
	case "+":
		fmt.Println(N1 + N2)
	case "-":
		fmt.Println(N1 - N2)
	case "*":
		fmt.Println(N1 * N2)
	case "/":
		// if N2 != 0 {
		// 	if N1 > N2 {
				if N2!=0{
					fmt.Printf("Result:%.2f\n",N1/N2)
		// 	} else {
		// 		fmt.Println(float64(N2/N1))
		// 	}	fmt.Printf("Result:%.2f\n",N1/N2)
			
		} else {
			fmt.Println("Cannot Devide with Number zero")
		}
	}
}
func main() {
	var N1 float64
	var N2 float64
	fmt.Println("Enter 1st number ")
	fmt.Scan(&N1)
	fmt.Println("Enter 2nd number")
	fmt.Scan(&N2)
	fmt.Println("Enter operation (+, -, *, /):")
	var operation string
	fmt.Scan(&operation)
	CliApp(N1, N2, operation)
}
