package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("The Title of the Book is ",b.title)
}

type count int

func (c count)String()string{
	return  fmt.Sprint("This is the number",strconv.Itoa(int(c)))
}
func main() {

	x:=book{
		title: "Rich Dad Poor Dad",
	}
	var y count =23

	log.Println(x)
	log.Println(y)
	log.Println("my name is aman")
}