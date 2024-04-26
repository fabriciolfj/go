package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	value := div(12, 6)
	fmt.Println(value)

	fmt.Println(add(3))
	fmt.Println(add(3, 4, 5, 6))

	result, rest, err := divAndRemainder(8, 3)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result, rest)
}

func div(num1, num2 int) int {
	if num1 == 0 || num2 == 0 {
		return 0
	}

	return num1 / num2
}

func add(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))

	for _, i := range vals {
		out = append(out, base+i)
	}

	return out
}

func divAndRemainder(num, demon int) (result int, rest int, err error) {
	if demon == 0 {
		err := errors.New("cannot divide by zero")
		return result, rest, err
	}

	result, rest = num/demon, num%demon
	return result, rest, nil
}
