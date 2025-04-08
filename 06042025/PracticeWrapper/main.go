package main

import "fmt"

type Devices interface {
	TurnOn() 
	TurnOff()
}

type Remote struct {
	Device string
}

func (r Remote) TurnOn()  {
	 fmt.Println("Remote is turning on of  ", r.Device)
}
func (r Remote) TurnOff()  {
	fmt.Println("Remote is turning off of  ", r.Device)
}

type Car struct {
	Brand string
}

func (c Car) TurnOn()  {
	 fmt.Println("Car is staring on of brand ", c.Brand)
}

func (c Car) TurnOff()  {
	 fmt.Println("Car is staring off of brand ", c.Brand)
}
func PrintDevice(d Devices) {
	d.TurnOn()
	d.TurnOff()
}

func main() {
	R1 := Remote{
		Device: "Tv",
	}
	R2 := Remote{
		Device: "Speaker",
	}
	Car1 := Car{
		Brand: "Toyota",
	}
	Car2 := Car{
		Brand: "Mahindra",
	}
	PrintDevice(R1)
	PrintDevice(R2)
	PrintDevice(Car1)
	PrintDevice(Car2)
}
