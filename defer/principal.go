package main

import "fmt"

func deferEample() int {
	a := 10

	defer func(val int) {
		fmt.Println("first:", val)
	}(a)

	a = 20

	defer func(val int) {
		fmt.Println("second", val)
	}(a)

	a = 30

	fmt.Println("exiting:", a)
	return a
}

func main() {
	deferEample()
}
