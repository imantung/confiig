package gofig

import "os"

type Handler interface {
	GetValue(name string) (val string, err error)
}

type EnvHandler struct{}

func NewEnvHandler() EnvHandler {
	return EnvHandler{}
}

func (h EnvHandler) GetValue(name string) (val string, err error) {
	val = os.Getenv(name)
	return
}
