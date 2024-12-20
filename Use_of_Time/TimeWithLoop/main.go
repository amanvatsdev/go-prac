package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a ticker that ticks every second
	ticker := time.Tick(1 * time.Second)

	// Loop to print the current time every second for 5 seconds
	for i := 1; i <= 5; i++ {
		// Wait for the next tick
		<-ticker

		// Print the current time
		fmt.Println("Current time:", time.Now().Format("15:04:05"))
	}
}

/*
2006=Year
01=Month
02=Day
03=2 Digit Time(12 Hour time)
15=2 Digit Time(24 Hour time)
04=Minutes
05=Seconds
Mon=To Specify The Day

*/
