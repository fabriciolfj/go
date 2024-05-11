package main

import (
	"fmt"
)

type Operator[T any] interface {
	Sum(T) T
}

type Calc struct {
	val1 int
	val2 int
}

func (c Calc) Sum(value int) int {
	return c.val1 + c.val2*value
}

func Execute[T any](op Operator[T], value T) {
	result := op.Sum(value)
	fmt.Println(result)
}

func main() {
	c := Calc{
		val1: 8,
		val2: 7,
	}

	Execute(c, 19)
}
