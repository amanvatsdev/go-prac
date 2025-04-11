package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("aman.go")
	if err !=nil{
		fmt.Println("error",err)
	}
	defer f.Close()
	s:=[]byte(`Hello World ,
	 My name is Aman
	 i am practicing go programming language`)

	 _,err=f.Write(s)
	 if err!=nil{
		log.Println("Error",err)
	 }
}