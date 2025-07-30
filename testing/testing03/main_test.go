package main

import (
	"log"
	"testing"
)

func TestIntro(t *testing.T){
	got:=Intro("Aman","Annu")
	want:="The name of students are Aman and Annu"
	if got != want{
		log.Fatalf("Error : got -%s want- %s",got,want)
	}
}
