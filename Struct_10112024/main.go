package main

import "fmt"

type Candidate struct {
	Name          string
	Age           int
	FathersName   string
	Aadhar_number int
	DOB           int
	Alive_Live    bool
	Mothers_Name  string
}

func main() {
	Person1 := Candidate{
		Name:          "Aman",
		Age:           25,
		DOB:           8_3_2004,
		Aadhar_number: 976902052900,
		Mothers_Name:  "R",
	}
	Peroson2 := Candidate{
		Mothers_Name: "hhh",
		Name:         "Rahul",
		DOB:          9_12_02,
	}

	Users_Data := []Candidate{Person1, Peroson2}
	// if Candidate.Aadhar_number.Users_Data !=0{
	// 	Candidate.Aadhar_number=false
	// }

	fmt.Printf("Here Is The Data Of Candidates \n%v", Users_Data)

	/* Type 2:==
	   func main() {
	   	Person1 := Candidate{
	   		Name:          "Aman",
	   		Age:           25,
	   		DOB:           8_3_2004,
	   		// Aadhar_number: 976902052900,
	   		Mothers_Name:  "R",
	   	}
	   	if Person1.Aadhar_number != 0 {
	   		Person1.Alive_Live = true
	   	} else {
	   		Person1.Alive_Live = false
	   	}
	   	fmt.Println(Person1)
	   }
	*/
	// .
	// .
	/*Practice 1:--
	  func main() {

	Printf()	Person1 := Candidate{
	  		Name:          "XYZ",
	  		Age:           25,
	  		DOB:           8_3_2004,
	  		Aadhar_number: 976902052900,
	  		Mothers_Name:  "R",
	  	}
	  	fmt.Println(Person1)
	  	Person1.Name = "Sinu"

	  	Person1.Age = 45
	  	fmt.Println(Person1)

	  }
	*/
}
