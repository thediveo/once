# `once`

[![Go Reference](https://pkg.go.dev/badge/github.com/thediveo/once.svg)](https://pkg.go.dev/github.com/thediveo/once)
![GitHub](https://img.shields.io/github/license/thediveo/once)
![build and test](https://github.com/TheDiveO/once/workflows/build%20and%20test/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/thediveo/once)](https://goreportcard.com/report/github.com/thediveo/once)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)

This tiny Go module improves your testing feng shui tremendiously by executing a
function only once that might need to be triggered from _multiple_ call sites
and not just from a _single_ `defer` call site.

Type `Oncer` works like `sync.Once`, except that the function to execute once is
specified only, _erm_, once when creating a `Oncer`.


#### Before

```go
rsc, _ := somemod.Open()
// Make sure to always clean up, but only once as
// somemod panics on multiple Close()s.
var once sync.Once
defer once.Do(func() { rsc.Close() })
// Do test something
// ...
once.Do(func() { rsc.Close() })
// Do test something else then f has been closed
// ...
```

#### After

```go
// You might want to dot-import for convenience.
import . "github.com/thediveo/once"

rsc, _ := somemod.Open()
// Make sure to always clean up, but only once as
// somemod panics on multiple Close()s.
once := Once(func() { rsc.Close() })
defer once.Do()
// Do test something
// ...
once.Do()
// Do test something else then f has been closed
// ...
```

## Go Version Support

`once` supports versions of Go that are noted by the [Go release
policy](https://golang.org/doc/devel/release.html#policy), that is, major
versions _N_ and _N_-1 (where _N_ is the current major version).

## Copyright and License

`once` is Copyright 2023 Harald Albrecht, and licensed under the Apache
License, Version 2.0.
