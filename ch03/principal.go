package main

import (
	"fmt"
	"slices"
)

func main() {
	/*arrays
	var x = [12]int{1, 5, 7, 10: 8}

	fmt.Println(x)

	var z = [...]int{1, 2, 3}
	fmt.Println(z[1])
	fmt.Println(len(z))

	//slices
	var y = []int{1, 2, 4}
	fmt.Println(y)
	fmt.Println(y[2])

	y[2] = 9
	fmt.Println(y)

	var t []int
	fmt.Println(t)

	fmt.Println(t == nil)
	*/
	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	//s := []string{"a", "b", "c"}

	fmt.Println(slices.Equal(x, y))
	fmt.Println(slices.Equal(x, z))
	//fmt.Println(slices.EqualFunc(x, s))

	x = append(x, 11)
	fmt.Println(x)

	x = append(x, y...)
	fmt.Println(x)

	p := [][]int{{1, 1}, {2, 3}}
	fmt.Println(p)
}
