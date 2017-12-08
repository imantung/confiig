package confiig

type Configurable interface {
	Register(name string) *Parameter
	Get(name string) (val Value, gErr GofigError)
}

type Config struct {
	paramMap map[string]*Parameter
	handler  Handler
}

func NewConfig(handler Handler) *Config {
	return &Config{
		paramMap: make(map[string]*Parameter),
		handler:  handler,
	}
}

func NewEnvConfig() *Config {
	return NewConfig(NewEnvHandler())
}

func (c *Config) SetHandler(handler Handler) {
	c.handler = handler
}

func (c *Config) Register(name string) (param *Parameter) {
	param = &Parameter{name: name}
	c.paramMap[name] = param
	return
}

func (c *Config) Param(name string) (param *Parameter) {
	return c.paramMap[name]
}

func (c *Config) Exist(name string) bool {
	_, ok := c.paramMap[name]
	return ok
}

func (c *Config) Get(name string) (val Value, gErr GofigError) {
	param, ok := c.paramMap[name]
	if !ok {
		gErr = MissingParamError(name)
		return
	}

	val, gErr = param.GetValue(c.handler)
	return
}

func (c *Config) Handler() Handler {
	return c.handler
}
