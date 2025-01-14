// package main

// import (
// 	"fmt"
// )

// type Grades struct {
// 	Maths int
// 	Hindi int
// 	Sci   int
// }
// type Address struct {
// 	StreetNo int
// 	City     string
// }
// type StudentInfo struct {
// 	Name       string
// 	Age        int
// 	Class      int
// 	FatherName string
// 	Grades     Grades
// 	Address    Address
// }

// func main() {
// 	fmt.Println("Type the Name of Student in terminal of which you want the data ")

// 	Students := []StudentInfo{
// 		{
// 			Name:       "Aman Kumar",
// 			Age:        15,
// 			Class:      10,
// 			FatherName: "Rajesh Kumar",
// 			Grades:     Grades{Maths: 85, Hindi: 78, Sci: 92},
// 			Address:    Address{StreetNo: 12, City: "Delhi"},
// 		},
// 		{
// 			Name:       "Neha Sharma",
// 			Age:        14,
// 			Class:      9,
// 			FatherName: "Ramesh Sharma",
// 			Grades:     Grades{Maths: 72, Hindi: 88, Sci: 80},
// 			Address:    Address{StreetNo: 7, City: "Mumbai"},
// 		},
// 		{
// 			Name:       "Rahul Verma",
// 			Age:        16,
// 			Class:      11,
// 			FatherName: "Suresh Verma",
// 			Grades:     Grades{Maths: 90, Hindi: 75, Sci: 85},
// 			Address:    Address{StreetNo: 5, City: "Kolkata"},
// 		},
// 		{
// 			Name:       "Priya Singh",
// 			Age:        15,
// 			Class:      10,
// 			FatherName: "Vinod Singh",
// 			Grades:     Grades{Maths: 88, Hindi: 82, Sci: 91},
// 			Address:    Address{StreetNo: 8, City: "Chennai"},
// 		},
// 		{
// 			Name:       "Rohit Das",
// 			Age:        14,
// 			Class:      9,
// 			FatherName: "Anil Das",
// 			Grades:     Grades{Maths: 76, Hindi: 84, Sci: 78},
// 			Address:    Address{StreetNo: 11, City: "Hyderabad"},
// 		},
// 		{
// 			Name:       "Sneha Patel",
// 			Age:        13,
// 			Class:      8,
// 			FatherName: "Amit Patel",
// 			Grades:     Grades{Maths: 79, Hindi: 80, Sci: 82},
// 			Address:    Address{StreetNo: 6, City: "Ahmedabad"},
// 		},
// 		{
// 			Name:       "Karan Mehta",
// 			Age:        16,
// 			Class:      11,
// 			FatherName: "Dinesh Mehta",
// 			Grades:     Grades{Maths: 85, Hindi: 78, Sci: 88},
// 			Address:    Address{StreetNo: 3, City: "Pune"},
// 		},
// 		{
// 			Name:       "Anjali Roy",
// 			Age:        15,
// 			Class:      10,
// 			FatherName: "Sunil Roy",
// 			Grades:     Grades{Maths: 92, Hindi: 89, Sci: 95},
// 			Address:    Address{StreetNo: 4, City: "Bangalore"},
// 		},
// 		{
// 			Name:       "Manish Gupta",
// 			Age:        14,
// 			Class:      9,
// 			FatherName: "Ashok Gupta",
// 			Grades:     Grades{Maths: 70, Hindi: 75, Sci: 72},
// 			Address:    Address{StreetNo: 10, City: "Lucknow"},
// 		},
// 		{
// 			Name:       "Pooja Yadav",
// 			Age:        13,
// 			Class:      8,
// 			FatherName: "Mahesh Yadav",
// 			Grades:     Grades{Maths: 82, Hindi: 85, Sci: 88},
// 			Address:    Address{StreetNo: 9, City: "Jaipur"},
// 		},
// 	}

// 	var FirstName, LastName string
// 	fmt.Println("Type the Name of student:")
// 	fmt.Scan(&FirstName)
// 	fmt.Scan(&LastName)
// 	Name := FirstName + " " + LastName
// 	// fmt.Scanf("%[^\n]", &Name)

// 	found := false
// 	for _, student := range Students {
// 		if student.Name == Name {
// 			fmt.Println("\nFull Data of " + student.Name + ":")
// 			fmt.Printf("Name: %s\n", student.Name)
// 			fmt.Printf("Age: %d\n", student.Age)
// 			fmt.Printf("Class: %d\n", student.Class)
// 			fmt.Printf("Father's Name: %s\n", student.FatherName)
// 			fmt.Printf("Grades: Maths=%d, Hindi=%d, Sci=%d\n", student.Grades.Maths, student.Grades.Hindi, student.Grades.Sci)
// 			found = true
// 			break
// 		}

// 	}
// 	if !found {
// 		fmt.Println("Student with Name", Name, "Not Found.")
// 	}
// }

// //i need to make a programm including data of students in school with their basic details including name class father nameffun
