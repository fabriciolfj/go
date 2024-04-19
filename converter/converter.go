package main

import "fmt"

func main() {
	x := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//y := x[2:]
	//z := x[:2]

	x[0] = 10

	//fmt.Println(x)
	//fmt.Println(y)
	//fmt.Println(z)
	fmt.Println(x[2:7])
}
