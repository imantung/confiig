# Confiig


Confiig separates between accessor and how to actually get the value which is offer idiomatic and clean way to setup project configuration. It's possible to make your own implementation without changing the library itself.

By default it's using env variable as [12 factor apps](https://12factor.net/) suggestion. 

### Register
```go
confiig.Register("name")
confiig.Register("name").Describe("description").Default("default value")
confiig.Register("name").Describe("description").Default(123)
```

### Get
```go
v, err := confiig.Get("name")
i := val.AsInt()
i, err = confiig.GetInt("name")
s, _ := config.GetString("name")
```

### Handler

```go
type CustomConfigHandler struct {
	ConfigMap map[string]string
}

func (h CustomConfigHandler) Handle(name string) (val string, err error) {
	for key, value := range h.ConfigMap {
		if strings.ToUpper(key) == strings.ToUpper(name) {
			val = value
			return
		}
	}

	err = fmt.Errorf("Missing %s on CustomConfig", name)
	return
}

func init(){
  handler := &CustomConfigHandler{ConfigMap: map[string]string{
		"host": "localhost",
		"port": "3306",
	}}
  
  confiig.SetHandler(handler)
  confiig.Register("host")
  
  host, _ := confiig.GetString("host")
  fmt.Println(host)
}
```


### As local variable

```go
conf := confiig.NewEnvConfig()
conf.Register("host")

s, _ := conf.GetString("host")
fmt.Println(s)
```
