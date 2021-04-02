/*
Package sponge implements a sponge function

Usage:
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
*/
package sponge
