package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("==================")

	for i := 0; i < 10; {
		if i%2 == 0 {
			i++
		} else {
			i += 2
		}

		fmt.Println(i)
	}

	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}
}
