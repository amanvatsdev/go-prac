package Currenttime

import (
	"fmt"
	"time"
)

func GetTheCurrentTime() {
	CurrentTime := time.Now()
	fmt.Println(CurrentTime.Format("2006,01,02  15:04:05"))
}
