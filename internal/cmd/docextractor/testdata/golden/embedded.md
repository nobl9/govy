# Embedded examples

[//]: # (embed: examples/full.go)

```go
package examples

const fullExample = "full"
```

[//]: # (embed: examples/named.go#ExampleNamed)

```go
// ExampleNamed demonstrates named extraction.
func ExampleNamed() {
	println("named")
}
```

[//]: # (embed: ExampleUnique)

```go
// ExampleUnique demonstrates bare lookup.
func ExampleUnique() {
	println("unique")
}
```
