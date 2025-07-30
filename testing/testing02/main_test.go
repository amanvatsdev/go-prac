package main

import "testing"

func TestAdd(t *testing.T){
	got:=Add(4,5)
	want:=9
	if got!=want{
		t.Errorf("Error: want %d got %d",want,got)
	}
}
