package main

import "fmt"

func returnErr(value bool) error {
	if value {
		return fmt.Errorf("deu ruim")
	}

	return nil
}

func main() {
	err := returnErr(false)
	if err != nil {
		fmt.Println(err)
	}

	err = returnErr(true)
	fmt.Println(err.Error())

}
