package main

import (
	Currenttime "aman/CurrentTime"
	"aman/ParseTime"
	"fmt"
	"time"
)

var name string = "Vikash"

func main() {
	Currenttime.GetTheCurrentTime()
	for {
		time.Sleep(3 * time.Second)
		fmt.Println(name)

		ParseTime.ParseTime()

	}

}
