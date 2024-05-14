package main

import (
	"fmt"
)

func dividirValures(a int, b int) (int, int, error) {
	if a == 0 {
		//return 0, 0, errors.New("valor 0 nao e divisivel") ou
		return 0, 0, fmt.Errorf("valor %d nao e divisivel", a)
	}

	return a / b, a % b, nil
}

func main() {
	a, b, error := dividirValures(0, 3)

	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println(a, b)
}
