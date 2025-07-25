package main

import (
	"errors"
	"fmt"
)

type NotImplementedError struct{}

func (*NotImplementedError) Error() string {
	return "internal not implemented"
}

func Translate1(err error) error {
	if (err == &NotImplementedError{}) {
		return errors.ErrUnsupported
	}

	return nil
}

func Translate2(err error) error {
	if (err == &NotImplementedError{}) {
		return nil
	}

	return err
}

func main() {
	fmt.Printf("translate1: %v\n", Translate1(DoWork()))
	fmt.Printf("translate2: %v\n", Translate2(DoWork()))
}

func DoWork() error {
	return &NotImplementedError{}
}
