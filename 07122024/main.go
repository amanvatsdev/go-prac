package main

import "fmt"

func main() {
	var marks int
	fmt.Println("Plz Type Your Total marks")
	fmt.Scanln(&marks)
	// var Grade string
	switch {
	case marks >80:
		fmt.Println(`Your Grade is :A
        Well done Keep it up`)
	case marks > 60:
		fmt.Println(`Your Grade is :B
        Well done Keep it Up`)
	case marks >400:
		fmt.Println(`Your Grade is :C
        You Need To Improve`)
	case marks >=0:
		fmt.Println(`Your Grade is :D
        You are Fail`)
	}

}
