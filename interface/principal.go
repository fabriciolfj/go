package main

import (
	"fmt"
)

type LogicProvider struct {
	describe string
}

func (lp LogicProvider) Process(data string) string {
	lp.describe = data
	return "step1"
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() string {
	return c.L.Process("client1")
}

func main() {
	c := Client{
		L: LogicProvider{},
	}

	fmt.Println(c.Program())

	var exemplo *string

	if exemplo == nil {
		fmt.Println("e nulo")
	}

	var exemplo2 string

	if exemplo2 == "" {
		fmt.Println("valor nao atribuido")
	}
}
