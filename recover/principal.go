package main

import "fmt"

func div(value int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()

	fmt.Println(60 / value)
}

func main() {
	for _, v := range []int{1, 2, 0, 5, 6} {
		div(v)
	}
}
