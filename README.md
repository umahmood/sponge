[![Go Reference](https://pkg.go.dev/badge/github.com/umahmood/sponge.svg)](https://pkg.go.dev/github.com/umahmood/sponge)

# Sponge

A Go library which implements a sponge function, from [Wikipedia](https://en.wikipedia.org/wiki/Sponge_function):

> In cryptography, a sponge function or sponge construction is any of a class of algorithms with finite internal state that take an input bit stream of any length and produce an output bit stream of any desired length. Sponge functions have both theoretical and practical uses. They can be used to model or implement many cryptographic primitives, including cryptographic hashes, message authentication codes, mask generation functions, stream ciphers, pseudo-random number generators, and authenticated encryption.

# Installation

```
$ go get github.com/umahmood/sponge@v1.0.0
```

See [VERSION](VERSION) or releases page for latest or specific version. 

# Usage

```
package main

import (
    "fmt"
    
    "github.com/umahmood/sponge"
)

func main() {
    s := sponge.New(nil)
    s.AbsorbByte(42)
    s.AbsorbBytes([]byte("Nothing will come of nothing."))
    bytes := s.Squeeze(14)
    fmt.Println(bytes)
}
```
By default the sponge will use SHA-256, to use another hash function:
```
s := sponge.New(md5.New())
```
Hash functions must implements the `hash.Hash` interface. see [here](https://golang.org/pkg/hash/#Hash).

# Documentation

[https://pkg.go.dev/github.com/umahmood/sponge](https://pkg.go.dev/github.com/umahmood/sponge)

# License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
