package confiig

import "fmt"

type Error interface {
	Param() *Parameter
	Error() string
}

type CommonError struct {
	param   *Parameter
	message string
}

// NewParamError
func NewError(param *Parameter, message string) Error {
	return &CommonError{param: param, message: message}
}

// NewError
func ConvertError(param *Parameter, err error) Error {
	return NewError(param, err.Error())
}

// NoValueError
func NoValueError(param *Parameter) Error {
	msg := fmt.Sprintf("Can't retrieve value from %s", param.Name())
	return NewError(param, msg)
}

// MissingParamError
func MissingParamError(name string) Error {
	msg := fmt.Sprintf("%s is not registered", name)
	return NewError(nil, msg)
}

func (e CommonError) Param() *Parameter {
	return e.param
}

func (e CommonError) Error() string {
	return e.message
}
