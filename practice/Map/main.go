package main

import "fmt"

type Address struct {
	StreetNo int
	City     string
	State    string
}
type EmployeeData struct {
	Name    string
	Salary  int
	Address Address
	Pan     string
}

func main() {
	/*Name , PhoneNo. , Pan , Salary , Address */
	//Adding Data Of Students
	Info := map[string]EmployeeData{
		"Aman": {
			Name:   "Aman",
			Salary: 500000,
			Address: Address{
				StreetNo: 54,
				City:     "Rohtak",
				State:    "Haryana",
			},
			Pan: "ABCPA1234E",
		},
		"Rahul": {
			Name:   "Rahul",
			Salary: 456000,
			Address: Address{
				StreetNo: 65,
				City:     "Jhajjar",
				State:    "Haryana",
			},
			Pan: "MHJPS5678L",
		},
		"Ajay": {
			Name:   "Ajay",
			Salary: 957000,
			Pan:    "WXTPL9876M",
			Address: Address{
				StreetNo: 75,
				City:     "Mumbai",
				State:    "Maharastra",
			},
		},
	}
	//Getting The Input From he Customer\
	// fmt.Println("Available Students:")
	// for name := range Info {
	// 	fmt.Println("-", name)
	// }
	fmt.Println("Tell me the Name of Employee of which you want the Data")
	var Name string
	fmt.Scan(&Name)
	Data, Exist := Info[Name]
	fmt.Println("Data : ", Data)
	fmt.Println("Exists : ", Exist)
	//Printing the Data
	if Exist {
		fmt.Print("Employee Data:")
		fmt.Printf("Name: %s\n", Data.Name)
		fmt.Printf("Salary: %d\n", Data.Salary)
		fmt.Printf("Address: %d, %s, %s\n", Data.Address.StreetNo, Data.Address.City, Data.Address.State)
		fmt.Printf("PAN: %s\n", Data.Pan)
	} else {
		fmt.Println("Employee not found in the records.")
	}

}
