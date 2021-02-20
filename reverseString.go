package main

import "fmt"
		"string"
func reverseString(s string){
			//word:=string.Fields(s)
			b := s.make([]string, 5)
			printSlice("b", b)
	
		}

func main() {
	var a [3]string
	fmt.Print("enter string")
	var input string
	for i := 0; i < 3; i++ {
		fmt.Scanf("%s", &input)
		a[i] = input
	}
	reverseString(a)
	
}
