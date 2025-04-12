	package main

	import (
		"fmt"
	
	)

	func processNumbers(numbers []int, callback func(int) (int,bool)) []int {
		var ProcessedNumbers []int

		for _,num:=range numbers{
			if val,ok:=callback(num);ok{
			ProcessedNumbers=append(ProcessedNumbers,val)
			}
		}
		return ProcessedNumbers
	}

	func main() {
		numbers := []int{2, 3, 4, 5}

		SquareEvenNumbers := func(num int) (int,bool) {
			
			if num%2 == 0	{
				return num*num,true
			}
			return 0,false
		}
		fmt.Println(processNumbers(numbers,SquareEvenNumbers))
	}


























// func filterNames (names []string,callback func(string)bool)[]string{
// 	var filterednames []string
// 	for _,name:=range names{
// 		if callback(name){
// 			filterednames=append(filterednames, name)
// 		}
// 	}
// 	return filterednames
// }
// func main() {
// 	names:=[]string{"Aman", "Ravi", "Ankit", "Aryan", "Dev"}

// 	LenthFilter:=func(x string)bool{
// 		return len(x)>=5
// 	}
// 	fmt.Println(filterNames(names,LenthFilter))
// }