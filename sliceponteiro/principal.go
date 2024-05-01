package main

import "fmt"

func addValues(list []int) {
	list = append(list, 4, 5, 6)
	list[0] = 9
	list[1] = 8
	list[2] = 7
}

func main() {
	list := make([]int, 3, 6)

	fmt.Println(list)

	addValues(list)
	fmt.Println(list)
}
