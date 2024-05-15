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
	var myError error
	if err != nil {
		if errors.As(err, &myError) {
			fmt.Println("that file doesn`t exist", myError)
		}
	}

}
