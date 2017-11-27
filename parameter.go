package gofig

import "fmt"

type Parameter struct {
	name         string
	description  string
	defaultValue interface{}
	handler      Handler
}

func (p *Parameter) GetValue() (val GofigValue, gErr GofigError) {
	str, err := p.handler.GetValue(p.name)
	if err != nil {
		gErr = ConvertError(p, err)
		return
	}

	if str == "" {
		if p.Mandatory() {
			gErr = NoValueError(p)
			return
		} else {
			str = p.DefaultValue()
		}
	}

	val = GofigValue(str)
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

func (p *Parameter) SetHandler(handler Handler) *Parameter {
	p.handler = handler
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

func (g *Parameter) Handler() Handler {
	return g.handler
}
