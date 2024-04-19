package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 2)

	num := copy(y, x)

	fmt.Println(y, num)
	fmt.Println(x[:])

}
