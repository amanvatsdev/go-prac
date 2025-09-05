package main

import (
	"fmt"
	"os"
)

func main() {

	//creation of  a file
	f, err := os.Create("text.txt")
	if err != nil {
		fmt.Println("error while creating file")
		return
	}
	defer f.Close() //closing the file

	//Entering data in the file
	data, err := f.Write([]byte("Hello world from Aman"))
	if err != nil {
		fmt.Println("error while entering the data :", err)
		return
	}
	fmt.Printf("Wrote %d bytes to file.\n", data)

	//opening the file again to read its data
	f2, err := os.Open("text.txt")
	if err != nil {
		fmt.Println("error while opening the file", err)
		return
	}
	defer f2.Close() // closing the file

	// creating a variable of []byte type
	m := make([]byte, 50)

	n, err := f2.Read(m)
	if err != nil {
		fmt.Println("error while reading the data", err)
		return
	}
	fmt.Println(n)
	fmt.Println(string(m[:n]))
}
