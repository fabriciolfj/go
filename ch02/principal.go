package main

import "fmt"

const test = "test"
const value = 1

func main() {
	//const test = "test" nao compila
	var (
		f int = 20
		g     = 4
	)

	x, z := 10, "fabricio"
	var y float64 = 21.99

	var sum = float64(x) + y
	var valueNew int = value

	fmt.Println(sum)
	fmt.Println(z)
	fmt.Println(f + g)
	fmt.Println(test)
	fmt.Println(valueNew)
}
