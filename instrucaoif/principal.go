package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if n := rand.Intn(10); n == 0 {
		fmt.Printf("valor 0, %d", n)
	} else if n > 5 {
		fmt.Printf("valor  maior que 5, %d", n)
	} else {
		fmt.Printf("valor muito pequeno, %d", n)
	}

	//fmt.Println(x) //nao compila
}
