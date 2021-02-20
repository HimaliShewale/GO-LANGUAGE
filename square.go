package main

import (
	"fmt"
	"math"
)

func main() {
	Square(4)
}

func Square(a float64) {
	x := math.Sqrt(a)
	fmt.Println("Square Root is", x)
}
