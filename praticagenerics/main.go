package main

import "fmt"

func Filter[T any](values []T, f func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range values {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

func main() {
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	f := func(v int) bool {
		return v%2 == 0
	}

	result := Filter(values, f)
	fmt.Println(result)
}
