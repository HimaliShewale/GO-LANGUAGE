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
	m:= map[string]person{
		p1.Name : p1,
		p2.Name : p2,
	}
	fmt.Println(m)
	for i,v := range m{
		fmt.Println(i,v)
	}
}
