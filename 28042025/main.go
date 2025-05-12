package main

import "fmt"

func BonusChecker(marks []int)int{
	var sum int
	for _, value := range marks {
		if value <= 50 {
			value += 5
		}
		sum += value
	}
	return sum
}
// 101159536042 
// 774359722774
func GradeCalculator(totalMarks int) (string,int){
	sum:=totalMarks/3
	var Grade string
	switch {
	case sum>=90:
		Grade="A"
	
	case sum>=50 :
		Grade="B"
	case sum>=33:
		Grade="C"
	}
	return Grade,sum
}

func main() {

	var x [3]int

	for index := 0; index < 3; index++ {

		fmt.Println("Plz enter your marks of subject ", index+1)
		fmt.Scanln(&x[index])
	}
	TotalMarks:=BonusChecker(x[:])
	fmt.Println(TotalMarks)
	g,m:=GradeCalculator(TotalMarks)
	fmt.Println("Your Grade is ",g,"and the average with bonus is",m)	
}
