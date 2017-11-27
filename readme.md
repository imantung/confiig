# Gofig (On Going)

Straight forward, lazy and singleton configuration for golang.


### Straight Forward

No one like to read documentation. We want a library that code example alone is enough to tell what we buy. Only 2 function required which is `register` to register our configuration and `get` to get the value of config. No matter whether it's environment variable, yaml, flag, remote config system or else, we want to keep the code still idiomatic. 

Register Function
```go
gofig.Register("name")
gofig.Register("name").Describe("description").Default("default value")
gofig.Register("name").Describe("description").Default(123)
gofig.Register("name").Handler(customHandler)
```

Get Function
```go
v, err := config.Get("name")
i := val.AsInt()
i, err = config.GetInt("name")

```

### Lazy

The value only retrieved until it's needed. 

### Singleton

Global access and only single instance. Hence this library is NOT for you:

- If you would like to keep your configuration as local variable 
- If you would like having more than 1 configuration on your apps.
