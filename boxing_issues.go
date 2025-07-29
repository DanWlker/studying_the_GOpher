package main

import (
	"fmt"
	"strings"
)

type MultiError struct {
	errs []string
}

func (m *MultiError) Add(err error) {
	m.errs = append(m.errs, err.Error())
}

func (m *MultiError) Error() string {
	return strings.Join(m.errs, ";")
}

func ReturnM() error {
	var m *MultiError

	return m
}

func ReturnNil() error {
	var m *MultiError

	if m != nil {
		return m
	}
	return nil
}

func main() {
	fmt.Println("ReturnM")
	if err := ReturnM(); err != nil {
		fmt.Println("err is not nil")
	}

	fmt.Println("=========")

	fmt.Println("ReturnNil")
	if err := ReturnNil(); err != nil {
		fmt.Println("err is not nil")
	}
}
