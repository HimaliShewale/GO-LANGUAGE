package main

import "fmt"

var x int = 42
var y string = "james"
var z bool = true

func main() {

	s := fmt.Sprintf("%v%v%v", x, y, z)
	fmt.Println(s)
	//fmt.Println(z)
}
