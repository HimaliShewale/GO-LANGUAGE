package main

import "fmt"

func main() {

	sum := 0
	var rem int

	fmt.Println("enter number")
	var num int
	fmt.Scanln(&num)
	new := num
	for num != 0 {
		rem = num % 10
		sum = sum*10 + rem
		num = num / 10
	}
	if sum == new {
		fmt.Println("number is palindrome")

	} else {
		fmt.Println("Not palindrome")
	}
}
