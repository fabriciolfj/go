package main

import "fmt"

type MyInt int

func main() {
	var i any
	var mine MyInt = 20
	i = mine

	i2, ok := i.(MyInt)
	if !ok {
		fmt.Errorf("valor invalido, espera-se um MyInt")
		return
	}

	fmt.Println(i2 + 1)

	doThings(i2)
	doThings("a")
}

func doThings(i any) {
	switch j := i.(type) {
	case nil:
		fmt.Println("variavel nula")
	case MyInt:
		fmt.Println("variavel e um MyInt", j)
	case string:
		fmt.Println("variavel é uma string", j)
	default:
		fmt.Println("impossivel determinar a variavel")
	}
}
