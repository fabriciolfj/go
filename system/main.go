package main

import (
	"fmt"

	"example.com/system/format"

	"example.com/system/math"
)

func main() {
	num := math.Double(2)
	output := format.Number(num)
	fmt.Println(output)
}
