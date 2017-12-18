package confiig

import "fmt"

type Parameter struct {
	name         string
	description  string
	defaultValue interface{}
}

func (p *Parameter) GetValue(handler Handler) (val Value, err Error) {
	str, err0 := handler.Handle(p.name)
	if err0 != nil {
		err = ConvertError(p, err0)
		return
	}

	if str == "" {
		if p.Mandatory() {
			err = NoValueError(p)
			return
		} else {
			str = p.DefaultValue()
		}
	}

	val = Value(str)
	return
}

func (p *Parameter) Describe(description string) *Parameter {
	p.description = description
	return p
}

func (p *Parameter) Default(defaultValue interface{}) *Parameter {
	p.defaultValue = defaultValue
	return p
}

func (p *Parameter) Mandatory() bool {
	return p.defaultValue == nil
}

func (p *Parameter) Name() string {
	return p.name
}

func (p *Parameter) Description() string {
	return p.description
}

func (g *Parameter) DefaultValue() string {
	return fmt.Sprint(g.defaultValue)
}
