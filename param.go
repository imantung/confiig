package gofig

import "fmt"

type Param struct {
	name         string
	description  string
	defaultValue interface{}
	handler      Handler
}

func (g *Param) Describe(description string) *Param {
	g.description = description
	return g
}

func (g *Param) Default(defaultValue interface{}) *Param {
	g.defaultValue = defaultValue
	return g
}

func (g *Param) SetHandler(handler Handler) *Param {
	g.handler = handler
	return g
}

func (g *Param) Mandatory() bool {
	return g.defaultValue == nil
}

func (g *Param) Name() string {
	return g.name
}

func (g *Param) Description() string {
	return g.description
}

func (g *Param) DefaultValue() string {
	return fmt.Sprint(g.defaultValue)
}

func (g *Param) Handler() Handler {
	return g.handler
}
