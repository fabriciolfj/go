package main

import "fmt"

type Integer interface {
	int | int64
}

func Convert[T1, T2 Integer](in T1) T2 {
	return T2(in)
}

func main() {
	var a int = 10
	b := Convert[int, int64](a)

	fmt.Println(b)
}
