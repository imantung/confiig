package confiig

import (
	"fmt"
	"strings"
	"testing"
)

type CustomConfigHandler struct {
	ConfigMap map[string]string
}

func (h CustomConfigHandler) Handle(name string) (val string, err error) {
	for key, value := range h.ConfigMap {
		if strings.ToUpper(key) == strings.ToUpper(name) {
			val = value
			return
		}
	}

	err = fmt.Errorf("Missing %s on CustomConfig", name)
	return
}

func TestCustomConfigHandler(t *testing.T) {

	handler := &CustomConfigHandler{ConfigMap: map[string]string{
		"Host": "localhost",
		"Port": "3306",
	}}
	config := NewConfig(handler)

	config.Register("HOST")
	config.Register("PORT")
	config.Register("NONAME")

	ifThenError(t, config.Handler() != handler, "handler is not set")

	val, err := config.Get("NONAME")
	ifThenError(t, err == nil, "expected error")
	ifThenError(t, val != "", "expected no value")
	ifThenError(t, err.Error() != "Missing NONAME on CustomConfig", "expect 'some error'")

	val, err = config.Get("PORT")
	ifThenError(t, err != nil, "expected no error")
	ifThenError(t, val != "3306", "expected 3306")

	val, err = config.Get("HOST")
	ifThenError(t, err != nil, "expected no error")
	ifThenError(t, val != "localhost", "expected localhost")

}
