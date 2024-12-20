package main

import "fmt"

type Information struct {
	Name        string
	Age         int
	live_status bool
	weight      float32
}

func main() {
	Candidate1 := Information{
		Name:   "Aman",
		Age:    42,
		weight: 80.5,
	}
	candidate2 := Information{
		Name:        "sinu",
		Age:         28,
		live_status: true,
	}
	Candidate1.live_status = true
	candidate2.weight = 75

	candidates := []Information{Candidate1, candidate2}

	for _, value := range candidates {
		fmt.Println(value)
	}
}
