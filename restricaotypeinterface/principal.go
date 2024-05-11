package main

import (
	"errors"
	"fmt"
)

type Integer interface {
	~int | ~int64
}

func divAndRemainder[T Integer](num, denon T) (T, T, error) {
	if denon == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	return num / denon, num % denon, nil
}

func main() {
	type MyInt int
	var my1 MyInt = 10
	var my2 MyInt = 10

	a, b, err := divAndRemainder(my1, my2)

	fmt.Println(a, b, err)
}
