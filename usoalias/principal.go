package main

import "fmt"

type Foo struct {
	X int
}

func (f Foo) hello() string {
	return "hello"
}

func (f Foo) goodbye() string {
	return "goodbye"
}

type Bar = Foo

func main() {
	result := Bar{
		X: 10,
	}

	fmt.Println(result)
	fmt.Println(result.goodbye())
}
