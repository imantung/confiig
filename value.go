package gofig

import "strconv"

type GofigValue string

func (v GofigValue) AsString() string {
	return string(v)
}

func (v GofigValue) AsInt() int {
	i, _ := strconv.Atoi(v.AsString())
	return i
}

func (v GofigValue) AsBool() bool {
	b, _ := strconv.ParseBool(v.AsString())
	return b
}

func (v GofigValue) AsFloat64() float64 {
	f, _ := strconv.ParseFloat(v.AsString(), 64)
	return f
}

func (v GofigValue) AsInt64() int64 {
	i, _ := strconv.ParseInt(v.AsString(), 10, 64)
	return i
}

func (v GofigValue) AsUint64() uint64 {
	u, _ := strconv.ParseUint(v.AsString(), 10, 64)
	return u
}
