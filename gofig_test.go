package gofig

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"
)

type AlwaysErrorHandler struct {
	errorMessage string
}

func (h AlwaysErrorHandler) GetValue(name string) (val string, err error) {
	err = fmt.Errorf(h.errorMessage)
	return
}

func TestSharedConfig(t *testing.T) {
	ifThenError(t, SharedConfig() != sharedConfig, "ShareConfig is retrieve share configurable")
}

func TestRegister(t *testing.T) {
	name := "TEST"
	param := Register(name).Describe("TEST_DESCRIPTION").Default(999)

	ifThenError(t, param.Name() != name, "Name is not set")
	ifThenError(t, param.Description() != "TEST_DESCRIPTION", "Description is not set")
	ifThenError(t, param.DefaultValue() != "999", "Default value is not set")
	ifThenError(t, param.Mandatory(), "Default value is not set")
	ifThenError(t, !Exist(name), "TEST is exist")
	ifThenError(t, Param(name) != param, "TEST param is same with param")

}

func TestGetTypo(t *testing.T) {
	val, err := Get("TYPO")

	ifThenError(t, val != "", "expected value to nil")
	ifThenError(t, err == nil, "expected error")
	ifThenError(t, err.Error() != "TYPO is not registered", "expect 'TYPO is not registered'")
	ifThenError(t, err.Param() != nil, "expect param is nil")
}

func TestGetNoDefaultAndNoSet(t *testing.T) {
	testNoDefault := Register("TEST_NO_DEFAULT")
	val, err := Get("TEST_NO_DEFAULT")

	ifThenError(t, val != "", "expected value to nil")
	ifThenError(t, err == nil, "expected error")
	ifThenError(t, err.Error() != "Can't retrieve value from TEST_NO_DEFAULT", "expect 'Can't retrieve value from TEST_NO_DEFAULT'")
	ifThenError(t, err.Param() != testNoDefault, "expect param is testNoDefault")
}

func TestCustomHandler(t *testing.T) {
	handler := AlwaysErrorHandler{errorMessage: "some error message"}
	param := Register("TEST_HANDLER").SetHandler(handler)
	ifThenError(t, param.Handler() != handler, "message is not set")

	val, err := Get("TEST_HANDLER")
	ifThenError(t, err == nil, "expected error")
	ifThenError(t, val != "", "expected no value")
	ifThenError(t, err.Error() != "some error message", "expect 'some error'")
}

func TestEnvHandler(t *testing.T) {
	Register("TEST_ENV")
	os.Setenv("TEST_ENV", "test-value")

	val, err := Get("TEST_ENV")
	ifThenError(t, err != nil, "expected error")
	ifThenError(t, val != "test-value", "expected value is 'test-value'")
}

func TestGetInt(t *testing.T) {
	Register("TEST_INT").Default(999)

	val, err := Get("TEST_INT")
	ifThenError(t, err != nil, "err is nil")
	ifThenError(t, val.AsInt() != 999, "val is int")
	ifThenError(t, val.AsInt64() != 999, "val is int64")
	ifThenError(t, val.AsUint64() != 999, "val is uint64")
}

func TestGetBool(t *testing.T) {
	Register("TEST_BOOL").Default(true)

	val, err := Get("TEST_BOOL")
	ifThenError(t, err != nil, "err is nil")
	ifThenError(t, val != "true", "val is not true")
	ifThenError(t, val.AsBool() != true, "val is bool")
}

func TestGetFloat(t *testing.T) {
	Register("TEST_FLOAT").Default(3.14)

	val, err := Get("TEST_FLOAT")
	ifThenError(t, err != nil, "err is nil")
	ifThenError(t, val.AsFloat64() != 3.14, "val is bool")
}

func ifThenError(t *testing.T, condition bool, message string) {
	if condition {
		_, file, no, ok := runtime.Caller(1)
		if ok {
			simpleFileName := file[strings.LastIndex(file, "/")+1:]
			message = fmt.Sprintf("%s:%d: %s", simpleFileName, no, message)
		}

		fmt.Println("%s", message)
		t.Fail()
	}
}
