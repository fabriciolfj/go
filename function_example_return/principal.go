package main

import (
	"errors"
	"fmt"
)

func cal(val int, mult int) (int, func(), error) {
	if mult == 0 {
		return 0, func() {
			fmt.Println("valor nao calculado")
		}, errors.New("valor 0 nao permitido")
	}

	return val * mult, func() {
		fmt.Println("valor calculado")
	}, nil
}

func main() {
	r, fun, error := cal(1, 0)

	if error != nil {
		fmt.Println(error)
		return
	}

	defer fun()
	fmt.Println("resultado", r)

}
