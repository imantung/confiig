package gofig

import "fmt"

type Parameter struct {
	name         string
	description  string
	defaultValue interface{}
	handler      Handler
}

func (g *Parameter) Describe(description string) *Parameter {
	g.description = description
	return g
}

func (g *Parameter) Default(defaultValue interface{}) *Parameter {
	g.defaultValue = defaultValue
	return g
}

func (g *Parameter) SetHandler(handler Handler) *Parameter {
	g.handler = handler
	return g
}

func (g *Parameter) Mandatory() bool {
	return g.defaultValue == nil
}

func (g *Parameter) Name() string {
	return g.name
}

func (g *Parameter) Description() string {
	return g.description
}

func (g *Parameter) DefaultValue() string {
	return fmt.Sprint(g.defaultValue)
}

func (g *Parameter) Handler() Handler {
	return g.handler
}
