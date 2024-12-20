package main

import "fmt"

func main() {
	type person struct {
		Name    string
		Age     int
		Hobbies []string
	}
	Persons := []person{
		{"Aman", 24, []string{"swimming", "Dancing"}},
		{"Sanjeev", 28, []string{}},
		{"Balraj", 45, []string{"Skating"}},
	}
	for _, people := range Persons {
		fmt.Printf("Name:%v\nAge;%v\n", people.Name, people.Age)

		switch {
		case people.Age <= 12:
			fmt.Println("Category: Child")
		case people.Age <= 21:
			fmt.Println("Category: teen")
		case people.Age <= 55:
			fmt.Println("Category: adult")
		default:
			fmt.Println("Category:old age")

			if len(people.Hobbies) > 0 {
				fmt.Printf("Hobbies-%v", people.Hobbies)
			} else {
				fmt.Println("Hobbies:No Hobbies")
			}
		}
	}
}
