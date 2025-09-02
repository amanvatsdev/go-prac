package main

import "fmt"

type data struct{}

func (d data) Write(m []byte) (x int, err error) {
	fmt.Println("Writing Date: ",string(m))
	return len(m),nil
}

func main() {
var x data

x.Write([]byte("Hello World its me Aman"))

}