package confiig

import "os"

type Handler interface {
	Handle(name string) (val string, err error)
}

type EnvHandler struct{}

func NewEnvHandler() EnvHandler {
	return EnvHandler{}
}

func (h EnvHandler) Handle(name string) (val string, err error) {
	val = os.Getenv(name)
	return
}
