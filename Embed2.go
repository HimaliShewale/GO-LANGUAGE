package main

import (
	"fmt"
)
type Vehicle struct{
	doors int
	color string
}
type Truck struct{
	Vehicle
	fourWheel bool
}
type Sedan struct{
	Vehicle
	luxury bool
}

func main(){
	p1 := Truck{
		Vehicle:Vehicle{
			doors: 4,
			color : "brown",
		},
		fourWheel: false,
	} 
	p2 := Sedan{
		Vehicle:Vehicle{
			doors: 3,
			color : "Black",
		},
		luxury: true,
	} 
	p3 := Vehicle{
		doors: 10,
		color : "Non",
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p1.doors,p1.color,p1.fourWheel)
	fmt.Println(p2.doors,p2.color,p2.luxury)
	fmt.Println(p3.doors,p3.color)
}