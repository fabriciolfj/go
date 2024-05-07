package main

import "fmt"

var Status int

const (
	VALID = iota
	INVALID
	ERROR
)

func main() {
	fmt.Println(VALID, INVALID, ERROR)
}
