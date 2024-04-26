package main

import "fmt"

func main() {
	var a string
	print := func(j int) {
		fmt.Println(j)
	}

	for i := 0; i < 10; i++ {
		print(i)
	}

	a = ""
	fmt.Println(a)
}
