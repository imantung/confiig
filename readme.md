# Gofig (On Going)

Straight forward, lazy and singleton configuration for golang.


### Straight Forward

No one like to read documentation. We want a library that code example alone is enough to tell what we buy. Only 2 function required which is `register` to register our configuration and `get` to get the value of config. No matter whether it's environment variable, yaml, flag, remote config system or else, we want to keep the code still idiom. 

Register Function
```go
gofig.Register("name").Describe("describe").Default("default value")
gofig.Register("name").Describe("describe").Type(int).Default(123)
gofig.Register("name").Describe("describe").Source(customSource)
```

Get Function
```go
val, err := config.Get("name")
intVal, err := config.GetInt("name")
val2 := config.GetOrDie("name")
```

### Lazy

Lazy is from lazy initialization. The value only retrieved until it's needed. 

### Singleton

Singleton is characteristic of global access and only single instance. Hence this library is NOT to:

- If You would like to keep your configuration as local variable 
- If you would like having more than 1 configuration on your apps.
