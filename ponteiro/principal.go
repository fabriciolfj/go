package main

import "fmt"

func main() {
	var x int32 = 10
	pointerX := &x

	fmt.Println(x)
	fmt.Println(pointerX)

	if pointerX != nil {
		fmt.Println(*pointerX)
	}

	//primitivos nao tem endereco de memoria, sao criados em tempo de execucao, para pegar o endereco utilize uma variavel auxiliar
	var y string
	z := &y
	fmt.Println(z)
}
