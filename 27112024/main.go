package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	StreetNo int    `json:"street_no,omitempty"`
	Village  string `json:"village,omitempty"`
	District string `json:"district,omitempty"`
	State    string `json:"state,omitempty"`
}
type Marks struct {
	Maths    int `json:"maths,omitempty"`
	Hindi    int `json:"hindi,omitempty"`
	English  int `json:"english,omitempty"`
	Sanskrit int `json:"sanskrit,omitempty"`
}
type Degrees struct {
	Name            string `json:"colledge,omitempty"`
	University      string `json:"university,omitempty"`
	UniversityState string `json:"university_state,omitempty"`
}

type StudentsDetails struct {
	Name    string    `json:"name,omitempty"`
	Address Address   `json:"address,omitempty"`
	Marks   Marks     `json:"marks,omitempty"`
	Hobbies []string  `json:"hobbies,omitempty"`
	Degrees []Degrees `json:"degrees,omitempty"`
}

func main() {
	Student1 := StudentsDetails{
		Name: "Jaideep",
		Address: Address{
			StreetNo: 25,
			Village:  "Kathura",
			District: "Rohtak",
			State:    "Haryana",
		},
		Marks: Marks{
			Maths:    90,
			Hindi:    98,
			English:  95,
			Sanskrit: 100,
		},
		// Hobbies: []string{"Singing", "Dancing", "Coding"},
		// Degrees: []Degrees{
		// 	{
		// 		Name:            "Bachelors in Commerce",
		// 		University:      "Delhi University",
		// 		UniversityState: "Delhi",
		// 	},
		// 	{
		// 		Name:            "Masters in Accounts",
		// 		University:      "Delhi University",
		// 		UniversityState: "Delhi",
		// 	},
		// },
	}

	JsonCoverteData, err := json.MarshalIndent(Student1, "#>>", " ")
	if err != nil {
		fmt.Println("This is the Error:", err)
	}

	fmt.Println(string(JsonCoverteData))

}
