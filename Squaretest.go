package main

import (
	"testing"
)

func TestSquareRoot(t * testing.T){
	if Square(4) != 2{
		t.Error("Your Answer is Wrong")
	}
}