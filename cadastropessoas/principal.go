package main

import (
	"fmt"

	"com.github/cadastropessoas/entity"
)

func main() {
	p := entity.Person{
		Name: "Fabricio",
		Age:  39,
	}

	fmt.Println(p)
}
