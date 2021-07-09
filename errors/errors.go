package errors

import (
	"errors"
	"fmt"
)

type GenericError struct {
	errType error
	msg     string
}

func (e *GenericError) Error() string { return fmt.Sprintf("%v: %s", e.errType, e.msg) }

var ErrNotFound = errors.New("not found")
var ErrResourceExhausted = errors.New("resource exhausted")

func Errorf(typ error, msg string, args ...interface{}) error {
	ret := new(GenericError)
	ret.errType = typ
	ret.msg = fmt.Sprintf(msg, args...)
	return ret
}
