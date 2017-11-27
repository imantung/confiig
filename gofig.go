package gofig

var sharedConfig = NewConfig()

func Register(name string) (param *Parameter) {
	param = sharedConfig.Register(name)
	return
}

func Get(name string) (val GofigValue, gErr GofigError) {
	val, gErr = sharedConfig.Get(name)
	return
}

func Param(name string) *Parameter {
	return sharedConfig.Param(name)
}

func Exist(name string) bool {
	return sharedConfig.Exist(name)
}

func SharedConfig() *Config {
	return sharedConfig
}
