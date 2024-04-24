package main

import "fmt"

func main() {
	envs := []int{1, 2, 3, 4}

	for _, v := range envs {
		fmt.Println(v)
	}

	maps := map[string]bool{"fabricio": true, "lucas": false}
	for k, v := range maps {
		fmt.Println(k, v)
	}
}
