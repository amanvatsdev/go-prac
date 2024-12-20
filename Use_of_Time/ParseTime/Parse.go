package ParseTime

import (
	"fmt"
	"time"
)

func ParseTime() {
	/*Meaning of Parse is that sometimes we have time and date in the format of string
	To Convert that time into Go Lang Time format we use Parse*/

	//giving a random var of string type containing time in it

	Time := "12:04:45 Mon"
	//syntax of parse is ParseTiem ,err:=time.parsed("Time Format,String type var")
	ParsedTime, err := time.Parse("15:04:05 Mon", Time)
	if err!=nil{
		fmt.Println("Error Occured during Parsuing")
	}
	fmt.Println(ParsedTime.Format("15:04:03 Mon"))
	

}