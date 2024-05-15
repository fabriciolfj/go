package main

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}

	f.Close()
	return nil
}

func main() {
	err := fileChecker("teste.txt")
	if err != nil {
		fmt.Println(err)
		if wappedErr := errors.Unwrap(err); wappedErr != nil {
			fmt.Println(wappedErr)
		}
	}
}
