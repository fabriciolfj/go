package main

import "fmt"

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}

	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]

	return top, true
}

func test[T2 any](values map[int]T2, key int) {
	ok, v := values[key]
	fmt.Println(ok, v)
}

func main() {
	var intStack Stack[int]
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)

	for _, i := range intStack.vals {
		v, ok := intStack.Pop()
		fmt.Println("indice ", i)
		fmt.Println("==========")
		fmt.Println(ok, v)

	}

	values := map[int]string{
		1: "Fabricio",
		2: "Lucas",
		3: "Jacob",
	}

	test(values, 1)
}
