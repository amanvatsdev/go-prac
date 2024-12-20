package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	StreetNo int    `json:"street_numbers,omitempty"`
	City     string `json:"city_name,omitempty"`
	Region   string `json:"region,omitempty"`
}
type Student struct {
	Name    string   `json:"name,omitempty"`
	Age     int      `json:"age,omitempty"`
	DOB     string   `json:"date_of_birth,omitempty"`
	Address Address  `json:"address,ommitempty"`
	Hobbies []string `json:"hobbies,omitempty"`
	Grades  string   `json:"Grade,omitempty"`
}

func GetTheGrades(Maths, Hindi, Sci int) string {
	var Grades string
	Numbers := Maths + Hindi + Sci
	Average := Numbers / 3
	switch {
	case Average > 80:
		Grades = "A"
	case Average > 60:
		Grades = "B"
	case Average > 40:
		Grades = "C"
	case Average >= 0:
		Grades = "D"

	}
	return Grades
}

func main() {
	Students := []Student{
		{Name: "Aman",
			Age: 21,
			DOB: "08-03-2004",
			Address: Address{
				StreetNo: 25,
				City:     "Delhi",
				Region:   "India",
			},
			Grades: GetTheGrades(55, 66, 85)},
		{Name: "Ravi",
			Age: 55,
			DOB: "05-03-2004",
			Address: Address{
				StreetNo: 42,
				City:     "Bombay",
				Region:   "India",
			},

			Grades: GetTheGrades(65, 99, 78)},
		{Name: "Vikash",
			Age: 32,
			DOB: "07-03-1994",
			Address: Address{
				StreetNo: 52,
				City:     "Kolkata",
				Region:   "India",
			},
			Grades: GetTheGrades(82, 96, 95)},
	}
	Data, _ := json.MarshalIndent(Students, "", " ")
	fmt.Println(string(Data))

	var UnmarshlData []Student

	err := json.Unmarshal([]byte(Data), &UnmarshlData)
	if err != nil {
		fmt.Println("Unmarshlling Error")
		return
	}
	fmt.Println("Decooded Code:")
	for _, value := range UnmarshlData {
		fmt.Printf("%+v\n", value)
	}
}
