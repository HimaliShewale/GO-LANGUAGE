package main

import "fmt"

func main() {
	var arr [5]int
	var input int
	fmt.Print("Enter the numbers is array")
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &input) //input from user
		arr[i] = input
	}
	fmt.Print(arr)
}
