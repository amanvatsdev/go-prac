package main

import "fmt"

// var PanCardNumber string = "xyzbi7895l"
// var name string = "sinu"
// var dob string = "04/05/2000"
// var fatherName string = "Vinod"

type User struct {
	FirstName     string
	LastName      string
	Age           int
	DOB           string
	MobileNo      string
	PanCardNumber string
	IsAlive       bool
}

func main() {
	var u = []User{

		{
			FirstName: "Sinu",
			LastName:  "Pandit",
		}, {
			FirstName: "ABC",
			LastName:  "XYZ",
		},
	}
	// fmt.Println(u)

	u[1].MobileNo = "8222909303"

	// fmt.Println(u)

	u[1].PanCardNumber = "AXCVB1234J"
	

	for i, v := range u {
		if v.FirstName == "Sinu" {
			u[i].IsAlive = true
			 
		}
	}

	fmt.Println(u)

}
