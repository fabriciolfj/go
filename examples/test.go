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

	for i, v := range v {
		if i == 2 {
			continue
		}

		fmt.Println(v)
	}
}
