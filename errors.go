package main

import (
	"errors"
	"fmt"
	"os"
)

type MyError string

const (
	ErrConst = MyError("this is my test error")
)

func OopsNonwrapped() error {
	return ErrConst
}

func OopsExists() error {
	return wrap{ErrConst, os.ErrExist}
}

func OopsInvalid() error {
	return wrap{ErrConst, os.ErrInvalid}
}

func (e MyError) Error() string {
	return string(e)
}

type wrap struct {
	msg MyError
	err error
}

func (w wrap) Error() string {
	return fmt.Sprintf("%v: %v", w.msg, w.err)
}

func (w wrap) Is(target error) bool {
	return errors.Is(w.msg, target) || errors.Is(w.err, target)
}

func (w wrap) Unwrap() error {
	return w.err
}

func main() {
	var toto *MyError
	e1 := OopsInvalid()

	if errors.As(e1, &toto) {
		fmt.Println(toto)
	}
	fmt.Print(e1)
}
