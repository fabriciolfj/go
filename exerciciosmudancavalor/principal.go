package main

import "fmt"

type Person struct {
	name string
}

func modifyName(person *Person) {
	person.name = person.name + "teste"
}

func main() {
	p := Person{
		name: "fabricio",
	}

	modifyName(&p)

	fmt.Println(p)
}
