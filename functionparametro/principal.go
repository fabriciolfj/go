package main

import (
	"fmt"
	"sort"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobdaughter", 23},
		{"Fred", "Fredson", 18},
	}

	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].lastName < people[j].lastName
	})

	fmt.Println(people)

}
