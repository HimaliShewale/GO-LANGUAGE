package main

import "fmt"

func main() {

	count := 0
	fmt.Println("enter the number")
	var num int
	fmt.Scanln(&num)
	//fmt.Println("sfsdg")
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			count++
		}
	}
	if count == 2 {
		fmt.Printf("%d is a prime Number", num)
	} else {
		fmt.Printf("%d is not a prime Number", num)
	}
}
