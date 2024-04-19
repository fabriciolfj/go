package main

import (
	"fmt"
	"maps"
)

func main() {
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2

	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Lions"])

	totalWins["Lions"]++
	fmt.Println(totalWins["Lions"])

	values := map[string]int{}

	values["hello"] = 1
	values["world"] = 1

	v, ok := values["hello"]
	fmt.Println(v, ok)

	v, ok = values["world"]
	fmt.Println(v, ok)

	v, ok = values["goodbye"]
	fmt.Println(v, ok)

	delete(values, "hello")
	v, ok = values["hello"]
	fmt.Println(v, ok)

	clear(values)

	fmt.Println(values)

	m := map[string]int{"hello": 5, "world": 6}
	n := map[string]int{"hello": 5, "world": 6}

	fmt.Println(maps.Equal(n, m))
}
