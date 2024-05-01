package main

import "fmt"

func failUpdate(px *int) {
	fmt.Println(px)
	x2 := 20
	px = &x2
}

func update(px *int) {
	*px = 20 //referencia ao ponteiro da variavel externa passada para essa funcao
}

func main() {
	var x int
	x = 10

	fmt.Println(x)

	failUpdate(&x)
	fmt.Println(x)

	update(&x)
	fmt.Println(x)
}
