package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)
	a[3] = 2
	fmt.Println(a)
	fmt.Println(len(a))
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	var twoD [2][3][4]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 4; k++ {
				twoD[i][j][k] = i + j + k
			}
		}
	}
	fmt.Println("2d: ", twoD)
}
