package main

import "fmt"

func main() {
	//tipo, tamanho, capacidade
	x := make([]int, 0, 10)
	x = append(x, 1, 2, 3)

	fmt.Println(x)
}
