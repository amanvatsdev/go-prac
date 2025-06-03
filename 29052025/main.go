package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Marks  float64 `json:"marks"`
	Stream string  `json:"stream"`
}

// Another to encode JSON

func TakeInput() Student {

	// Use one function to take input
	var name string
	var age int
	var marks float64
	var stream string

	fmt.Printf("Enter Name:")
	fmt.Scanln(&name)

	fmt.Printf("Age:")
	fmt.Scanln(&age)

	fmt.Print("Marks:")
	fmt.Scanln(&marks)

	fmt.Print("Stream:")
	fmt.Scanln(&stream)

	fmt.Println("")

	student1 := Student{
		Name:   name,
		Age:    age,
		Marks:  marks,
		Stream: stream,
	}
	// fmt.Printf("Here is the data of student 1\n%v", student1)
	return student1
}

// Another function to calculate grade
func GradeCalculater(marks float64) {
	if marks >= 90 {
		fmt.Println("Grade=A")
	} else if marks >= 75 {
		fmt.Println("Grade=B")
	} else if marks >= 50 {
		fmt.Println("Grade=C")
	} else {
		fmt.Println("Grade=F")
	}
}

func EncodeJson(x Student) {
	json, err := json.Marshal(x)
	if err != nil {
		fmt.Println("json error")
		return
	}
	fmt.Printf("Student Info(Json):\n%v\n\n", string(json))
}

func main() {

	Student1 := TakeInput()
	EncodeJson(Student1)
	GradeCalculater(Student1.Marks)

}
