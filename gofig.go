package gofig

var paramMap = make(map[string]*Param)

func Register(name string) *Param {
	Param := &Param{
		name:    name,
		handler: NewEnvHandler(),
	}
	paramMap[name] = Param
	return Param
}

func Get(name string) (val GofigValue, gErr GofigError) {
	param, ok := paramMap[name]
	if !ok {
		gErr = MissingParamError(name)
		return
	}

	s, err := param.handler.GetValue(name)
	if err != nil {
		gErr = ConvertError(param, err)
		return
	}

	if s == "" {
		if param.Mandatory() {
			gErr = NoValueError(param)
			return
		} else {
			s = param.DefaultValue()
		}
	}

	val = GofigValue(s)
	return
}
