package main

import "fmt"

func main() {
	x := []string{}
	x = append(x, "fabricio", "suzana", "teste")

	fmt.Println(x)

	clear(x)
	fmt.Println(x, len(x))
}
