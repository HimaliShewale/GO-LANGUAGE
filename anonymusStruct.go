package main
import (
	"fmt"
)
func main() {
	//anonymus struct go
	p1 := struct {
		Name string
		Age  int
	}{
		Name: "Himali",
		Age:  25,
	}
	
	fmt.Println(p1)

	fmt.Println(p1.Name)

}
