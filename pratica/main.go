package main

import (
	"fmt"
	"strconv"
)

func multiply(a int) func(int) int {
	if a%2 == 0 {
		return func(a int) int {
			return a * 2
		}
	}

	return func(a int) int {
		return a * 3
	}

}

func main() {
	fun := multiply(2)

	result := fun(2)
	fmt.Printf(strconv.Itoa(result))
}
