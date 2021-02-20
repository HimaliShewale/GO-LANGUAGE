package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter number")
	fmt.Scanf("%d", &num)
	sum := 0
	rem := 0
	temp := num

	for num != 0 {
		rem = num % 10
		sum = sum + rem*rem*rem
		num = num / 10
	}
	if temp == sum {
		fmt.Print("number is armstrong")
	} else {
		fmt.Print("number is not armstrong")
	}
}
