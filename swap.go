package main

import "fmt"

func Swap(x string, y string) (string, string) {
	return y, x
}

func main() {

	a, b := Swap("hello", "world")
	fmt.Println(a, b)
}
