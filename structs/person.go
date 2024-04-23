package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func main() {
	p := person{
		"fabricio",
		39,
		"dog",
	}

	fmt.Println(p.name)

	type first struct {
		name string
		age  int
	}

	type second struct {
		name string
		age  int
	}

	f := first{name: "lucas", age: 18}
	g := first{name: "lucas", age: 18}

	fmt.Println(f == g)

	//e := second{name: "lucas", age: 18}

	//fmt.Println(f == e) not compile

	e := []first{}
	e = append(e, f, g)

	fmt.Println(e)
}
