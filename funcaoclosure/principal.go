package main

import "fmt"

func increment() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func main() {
	inc := increment()

	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}
