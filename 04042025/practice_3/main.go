package main

import "fmt"

type Device interface {
	TurnOn()
	TurnOff()
}

type TV struct {
	Brand string
}

func (t TV) TurnOn() {
	fmt.Println("The Tv is starting up...")
}
func (t TV)TurnOff(){
	fmt.Println("The Tv is turning off..")
}
type AC struct{
	Brand string
}
func (a AC) TurnOn() {
	fmt.Println("The Ac is starting up...")
}
func (a AC)TurnOff(){
	fmt.Println("The Ac is turning off..")
}

type MusicSystem struct{
	Brand string
}
func (m MusicSystem) TurnOn() {
	fmt.Println("The MusicSystem is starting up...")
}
func (m MusicSystem)TurnOff(){
	fmt.Println("The MusicSystem is turning off..")
}

func PrintDevice(d Device){
	d.TurnOn()
	d.TurnOff()
}
func main() {
	Tv1:=TV{
		Brand: "LG",
	}
	Ac1:=AC{
		Brand: "Goodrej",
	}
	Ms1:=MusicSystem{
		Brand: "Boat",
	}

	PrintDevice(Tv1)
	PrintDevice(Ac1)
	PrintDevice(Ms1)


}