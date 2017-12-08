package confiig

import "fmt"

type GofigError interface {
	Param() *Parameter
	Error() string
}

type GenericGofigError struct {
	param   *Parameter
	message string
}

// NewParamError
func NewError(param *Parameter, message string) GenericGofigError {
	return GenericGofigError{param: param, message: message}
}

// NewError
func ConvertError(param *Parameter, err error) GenericGofigError {
	return NewError(param, err.Error())
}

// NoValueError
func NoValueError(param *Parameter) GenericGofigError {
	msg := fmt.Sprintf("Can't retrieve value from %s", param.Name())
	return NewError(param, msg)
}

// MissingParamError
func MissingParamError(name string) GenericGofigError {
	msg := fmt.Sprintf("%s is not registered", name)
	return NewError(nil, msg)
}

func (e GenericGofigError) Param() *Parameter {
	return e.param
}

func (e GenericGofigError) Error() string {
	return e.message
}
