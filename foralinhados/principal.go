package main

import "fmt"

func main() {
	samples := []string{"hello", "apple"}
outer:
	for _, sample := range samples {
		for _, v := range sample {
			fmt.Println(string(v))

			if v == 'l' {
				continue outer
			}
		}

		fmt.Println()
	}
}
