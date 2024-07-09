package main

import (
	"fmt"
	"strconv"
)

type Geometry interface {
	Area() int
}

type Square struct {
	v int
}

type Rectangle struct {
	base   int
	height int
}

func (s *Square) Area() int {
	return s.v * s.v
}

func (r *Rectangle) Area() int {
	return r.base * r.height
}

func execute(g Geometry) int {
	return g.Area()
}

func check(g Geometry) {
	switch g.(type) {
	case *Square:
		fmt.Printf("square area")
	case *Rectangle:
		fmt.Println("rectangle area")
	default:
		fmt.Println("unknow type")
	}
}

func main() {
	square := Square{3}
	rectangle := Rectangle{4, 6}
	fmt.Println(strconv.Itoa(execute(&square)))
	fmt.Println(strconv.Itoa(execute(&rectangle)))

	check(&square)
}
