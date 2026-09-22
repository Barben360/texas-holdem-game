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
var ErrUnimplemented = errors.New("unimplemented")
var ErrBadRequest = errors.New("bad request")
var ErrInternal = errors.New("internal error")
var ErrUnknown = errors.New("unknown error")

func Errorf(typ error, msg string, args ...interface{}) error {
	ret := new(GenericError)
	ret.ErrType = typ
	ret.Msg = fmt.Sprintf(msg, args...)
	return ret
}

func ErrorfAppend(oldErr error, msg string, args ...interface{}) error {
	e, ok := oldErr.(*GenericError)
	if !ok {
		return Errorf(ErrUnknown, msg+" - "+e.Error(), args...)
	}
	return Errorf(e.ErrType, msg+" - "+e.Msg, args...)
}

func ErrorfUsingType(oldErr error, msg string, args ...interface{}) error {
	e, ok := oldErr.(*GenericError)
	if !ok {
		return Errorf(ErrUnknown, msg, args...)
	}
	return Errorf(e.ErrType, msg, args...)
}
