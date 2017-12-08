# Confiig (On Going)

Straight forward & Decoupled configuration for golang.

### Straight Forward

Register Function
```go
gofig.Register("name")
gofig.Register("name").Describe("description").Default("default value")
gofig.Register("name").Describe("description").Default(123)
```

Get Function
```go
v, err := gonfig.Get("name")
i := val.AsInt()
i, err = gonfig.GetInt("name")

```

### Decoupled

Confiig separate between setter/getter and the implementation. By default it's using env variable as 12 factor of apps suggstion. This library also help if you want to use your own custom like using redis or your own formatted config.
