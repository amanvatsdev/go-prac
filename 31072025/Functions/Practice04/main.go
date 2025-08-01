package main

import (
	"fmt"
	"time"
)

func main() {
timefunc(doWork)

}

func doWork(){
	for a:=0;a<2_000;a++{
		fmt.Println(a)
	}
}

func timefunc(f func()) {
	start:=time.Now()
	f()
	elapsed:=time.Since(start)
	fmt.Println(elapsed)

}
