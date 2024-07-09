package main

import "fmt"

func main() {
	a := make([]string, 10)

	a[0] = "fabricio"
	a[1] = "lucas"
	a[2] = "felicio"
	a[3] = "jacob"
	a[4] = "penna"

	for _, v := range a {
		fmt.Println(v)
	}

	v := []int{1, 2, 3, 4}
	z := []int{5, 6, 7, 8}

	p := make([]int, 0, 10)
	p = append(p, v...)
	p = append(p, z...)

	for i, q := range p {
		if i == 2 {
			continue
		}

		fmt.Println(q)
	}
}