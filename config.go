package gofig

type Configurable interface {
	Register(name string) *Parameter
	Get(name string) (val GofigValue, gErr GofigError)
}

type Config struct {
	paramMap map[string]*Parameter
}

func NewConfig() *Config {
	return &Config{
		paramMap: make(map[string]*Parameter),
	}
}

func (c Config) Register(name string) (param *Parameter) {
	param = &Parameter{
		name:    name,
		handler: NewEnvHandler(),
	}
	c.paramMap[name] = param
	return
}

func (c Config) Param(name string) (param *Parameter) {
	return c.paramMap[name]
}

func (c Config) Exist(name string) bool {
	_, ok := c.paramMap[name]
	return ok
}

func (c Config) Get(name string) (val GofigValue, gErr GofigError) {
	param, ok := c.paramMap[name]
	if !ok {
		gErr = MissingParamError(name)
		return
	}

	val, gErr = param.GetValue()
	return
}
