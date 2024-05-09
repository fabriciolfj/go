package main

import "fmt"

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}

	return r
}

func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}

	return r
}

func Filter[T any](s []T, f func(T) bool) []T {
	r := make([]T, 0)
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}

	return r
}

func main() {
	words := []string{"One", "Potato", "Two", "Potato"}
	filtered := Filter(words, func(s string) bool {
		return s != "Potato"
	})

	fmt.Println(filtered)
}
