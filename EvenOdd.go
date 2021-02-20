package main

import "fmt"

func main() {
	fmt.Println("Enter the number")
	var num int
	fmt.Scanln(&num)
	if num%2 == 0 {
		fmt.Printf("%d is Even Number", num)
	} else {
		fmt.Printf("%d is Odd Number", num)
	}

}
