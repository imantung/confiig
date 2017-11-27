package gofig

import "fmt"

type GofigError interface {
	Param() *Param
	Error() string
}

type GenericGofigError struct {
	param   *Param
	message string
}

// NewParamError
func NewError(param *Param, message string) GenericGofigError {
	return GenericGofigError{param: param, message: message}
}

// NewError
func ConvertError(param *Param, err error) GenericGofigError {
	return NewError(param, err.Error())
}

// NoValueError
func NoValueError(Param *Param) GenericGofigError {
	msg := fmt.Sprintf("Can't retrieve value from %s", Param.Name())
	return NewError(Param, msg)
}

// MissingParamError
func MissingParamError(name string) GenericGofigError {
	msg := fmt.Sprintf("%s is not registered", name)
	return NewError(nil, msg)
}

func (e GenericGofigError) Param() *Param {
	return e.param
}

func (e GenericGofigError) Error() string {
	return e.message
}
