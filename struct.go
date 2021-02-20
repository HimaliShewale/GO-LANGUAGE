package main
import (
	"fmt"
)

type person struct {
	Name string
	Age  int
	Ice []string
}

func main() {
	p1 := person{
		Name: "Himali",
		Age:  25,
		Ice : []string{"Choclate","Vvanilla",},
	}
	p2 := person{
		Name: "Rutuja",
		Age:  20,
		Ice : []string{"Vanilla","Stawberry",},
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.Name)
	fmt.Println(p2.Ice)
}
