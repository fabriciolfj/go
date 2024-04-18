package main

import "fmt"

func main() {
	//compartilha a mesma memoria
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	//pega o valor 1 2
	y := x[:2:5]

	fmt.Println(y)

	//coloca no final do tamanho, mas modifica o valor de x por é compartilhado, substituindo os valores nas respectivas posicoes
	y = append(y, 11, 12, 13)

	fmt.Println(y)
	fmt.Println(x) //como foi alterado o resultado agora e [1 2 11 12 13 6 7 8 9 10]
}
