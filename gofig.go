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

func GetInt(name string) (i int, gErr GofigError) {
	val, gErr := Get(name)
	i = val.AsInt()
	return
}

func GetInt64(name string) (i int64, gErr GofigError) {
	val, gErr := Get(name)
	i = val.AsInt64()
	return
}

func GetUint64(name string) (i uint64, gErr GofigError) {
	val, gErr := Get(name)
	i = val.AsUint64()
	return
}

func GetString(name string) (s string, gErr GofigError) {
	val, gErr := Get(name)
	s = val.AsString()
	return
}

func GetBool(name string) (b bool, gErr GofigError) {
	val, gErr := Get(name)
	b = val.AsBool()
	return
}

func GetFloat64(name string) (f float64, gErr GofigError) {
	val, gErr := Get(name)
	f = val.AsFloat64()
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
