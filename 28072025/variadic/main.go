package main

import (
	"fmt"
)

func main() {

a:=[]int{34,54,2,64,3,34,64,34,53,34}
b:=foo(a...)
fmt.Println(b)

c:=bar(a)
fmt.Println(c)

}

func foo (a ...int)int{
	sum:=0
	for _,v:= range a{
		sum += v
	}
	return sum
}

func bar (a []int)(sum int){
	for _,v:=range a{
		sum +=v
	}
	return
}