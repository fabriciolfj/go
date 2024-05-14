package main

import "fmt"

type Status int

const (
	INvalidLogin Status = iota + 1
	NotFound
)

func (se StatusError) Error() string {
	return se.Message
}

type StatusError struct {
	status  Status
	Message string
}

func GenerateErrorUseErrorVar(flag bool) error {
	var genError error
	if flag {
		genError = StatusError{
			status:  NotFound,
			Message: "",
		}
	}

	return genError
}

func main() {
	err := GenerateErrorUseErrorVar(true)
	fmt.Println(err != nil)

	err = GenerateErrorUseErrorVar(false)
	fmt.Println(err != nil)
}
