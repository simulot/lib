package myerrors

import (
	"errors"
	"fmt"
	"strings"
)

func Should(v interface{}, err error) interface{} {
	if err != nil {
		return nil
	}
	return v
}

type MyErrors struct {
	errs []error
}

func New() *MyErrors {
	return &MyErrors{}
}

func (e *MyErrors) Should(v interface{}, err error) interface{} {
	if err != nil {
		e.Check(err)
	}
	return v
}
func (e *MyErrors) Check(err error) error {
	if err == nil {
		return nil
	}
	var myErr *MyErrors
	if errors.As(err, &myErr) {
		e.errs = append(e.errs, myErr.errs...)
		return err
	}
	e.errs = append(e.errs, err)
	return err
}

func (e *MyErrors) IsNotEmpty(s string, err error) error {
	e.Check(err)
	return err
}

func (e *MyErrors) Err() error {
	if len(e.errs) > 0 {
		return e
	}
	return nil
}

func (e *MyErrors) Error() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("%d errors have been found:\n", len(e.errs)))
	for _, e := range e.errs {
		s.WriteString(e.Error())
		s.WriteString("\n")
	}
	return s.String()
}
