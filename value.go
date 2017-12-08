package confiig

import "strconv"

type Value string

func (v Value) AsString() string {
	return string(v)
}

func (v Value) AsInt() int {
	i, _ := strconv.Atoi(v.AsString())
	return i
}

func (v Value) AsBool() bool {
	b, _ := strconv.ParseBool(v.AsString())
	return b
}

func (v Value) AsFloat64() float64 {
	f, _ := strconv.ParseFloat(v.AsString(), 64)
	return f
}

func (v Value) AsInt64() int64 {
	i, _ := strconv.ParseInt(v.AsString(), 10, 64)
	return i
}

func (v Value) AsUint64() uint64 {
	u, _ := strconv.ParseUint(v.AsString(), 10, 64)
	return u
}
