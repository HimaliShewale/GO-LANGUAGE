package main

import (
	"fmt"
)

type People struct {
	Name string
	Age  int
}
type person struct {
	People
	flag bool
}

func main() {
	per1 := person{
		People: People{
			Name: "Om",
			Age:  15,
		},
		flag :true,
	}
	p2 := People{
		Name: "Rutuja",
		Age:  20,
	}
	fmt.Println(per1)
	fmt.Println(per1)
	fmt.Println(per1.Name)
	fmt.Println(p2.Age)
}
