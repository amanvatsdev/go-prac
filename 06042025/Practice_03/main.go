package main

import "fmt"

type Player interface {
	Play()
}

type Audioplayer struct {
	Devicename string
	Brand      string
}

func (a Audioplayer) Play() {
	fmt.Println(a.Devicename,"Started playing music of brand ",a.Brand)
}

type Videoplayer struct {
	Devicename string
	Brand      string
}
func (v Videoplayer) Play() {
	fmt.Println(v.Devicename,"Started playing music of brand ",v.Brand)
}
func printplayer(p Player){
	p.Play()
}
func main(){
	A1:=Audioplayer{
		Devicename: "Speaker",
		Brand: "JBL",
	}
	A2:=Audioplayer{
		Devicename: "Headphones",
		Brand: "Boat",
	}
	V1:=Videoplayer{
		Devicename: "Projector",
		Brand:"Lg",
	}
	V2:=Videoplayer{
		Devicename: "Monitor",
		Brand: "Lenovo",
	}
	printplayer(A1)
	printplayer(A2)
	printplayer(V1)
	printplayer(V2)
}
