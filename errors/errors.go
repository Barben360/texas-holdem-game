package errors

import (
	"errors"
	"fmt"
)

type GenericError struct {
	ErrType error
	Msg     string
}

func (e *GenericError) Error() string { return fmt.Sprintf("%v: %s", e.ErrType, e.Msg) }

var ErrNotFound = errors.New("not found")
var ErrResourceExhausted = errors.New("resource exhausted")
