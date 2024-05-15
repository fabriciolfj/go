package main

import (
	"errors"
	"fmt"
)

func doThing1() (string, error) {
	return "", errors.New("teste1")
}

func doThing2() (string, error) {
	return "", errors.New("teste2")
}

func doSomeThings() (_ string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("in doSomeThings: %w", err)
		}
	}()

	val1, err := doThing1()
	if err != nil {
		return val1, err
	}

	return doThing2()
}

func main() {
	var err error
	var result string
	result, err = doSomeThings()
	fmt.Println(err, result)
}
