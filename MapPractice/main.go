package main

import "fmt"

func main() {
	Salaries := map[string]int{
		"Alice":   50000,
		"BoB":     60000,
		"Charlie": 55000,
	}
	fmt.Println("Choose an option")
	fmt.Println(`
	1.Add a new employee with their salary.
	2.Update the salary of an existing employee.
	3.Retrieve the salary of a specific employee.
	4.Delete an employee's record.
	5.Calculate and display the total salary of all employees.`)
	var option int
	fmt.Scanln(&option)
	switch option {
	case 1:
		var Name string
		fmt.Print("Enter Name:")
		fmt.Scanln(&Name)
		var Salary int
		fmt.Print("Enter Salary:")
		fmt.Scanln(&Salary)
		Salaries[Name] = Salary
		fmt.Println("Employee Name Added:Name of student:", Name, "His Salary is :", Salary)
	case 2:
		fmt.Println("Enter The Name of Student")
		var Name string
		fmt.Scanln(&Name)
		_, Exist := Salaries[Name]
		if Exist {
			fmt.Println("Enter the Salary")
			var Salary int
			fmt.Scanln(&Salary)
			Salaries[Name] = Salary
			fmt.Println("Salary is Updated")
		} else {
			fmt.Println("Name Doesn;t Exist")
		}
	case 3:
		fmt.Println("Enter the Name ")
		var Name string
		fmt.Scanln(&Name)
		Salary, Exist := Salaries[Name]
		if Exist {
			fmt.Printf("Name:%s\tSalary:%d", Name, Salary)
		} else {
			fmt.Println("Name Doest exist in the list")
		}
	case 4:
		fmt.Println("Enter the Name of Student")
		var Name string
		fmt.Scanln(&Name)
		_, Exist := Salaries[Name]
		if Exist {
			delete(Salaries, Name)
			fmt.Println(Name, "This name was no longer in the list")
		} else {
			fmt.Println("This Name Was Not In The List")
		}
	case 5:

		if len(Salaries) == 0 {
			fmt.Println("No Member Added")
		} else {
			fmt.Println("Calculating Salary Please Wait")
			Sum := 0
			for _, Salary := range Salaries {
				Sum += Salary
			}
			fmt.Println("Total Salary of Employee", Sum)

		}
	default:
		fmt.Println("Please Select a proper option between 1 to 5 intergers only")
	}
}
