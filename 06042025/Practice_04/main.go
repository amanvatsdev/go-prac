package main

import "fmt"

type Appliance interface {
	Start()
	Stop()
}
type WashingMachine struct {
	Brand string
	Model string
}

func (w WashingMachine) Start() {
	fmt.Println(w.Brand, "is started of model", w.Model)
}
func (w WashingMachine) Stop() {
	fmt.Println(w.Brand, "is stopped of model", w.Model)
}

type Microwave struct {
	Brand string
	Model string
}

func (m Microwave) Start() {
	fmt.Println(m.Brand, "is started of model ", m.Model)
}
func (m Microwave) Stop() {
	fmt.Println(m.Brand, "is stopped of model", m.Model)
}

type Refrigerator struct {
	Brand string
	Model string
}

func (r Refrigerator) Start() {
	fmt.Println(r.Brand, "is started of model ", r.Model)
}
func (r Refrigerator) Stop() {
	fmt.Println(r.Brand, "is stopped of model", r.Model)
}
func ControlAppliance(ap Appliance) {
	ap.Start()
	ap.Stop()
}

func main(){
	WM1:=WashingMachine{
		Brand: "Whirlpool",
		Model: "LG GL-I292RPZL",
	}
	M1:=Microwave{
		Brand: "Voltas",
		Model:"MC28A5013AK",
	}
	R1:=Refrigerator{
		Brand: "Lg",
		Model: "GL-I292RPZL",
	}
	fmt.Println("------------------------------------")
	ControlAppliance(WM1)
	fmt.Println("------------------------------------")
	ControlAppliance(M1)
	fmt.Println("------------------------------------")
	ControlAppliance(R1)
}