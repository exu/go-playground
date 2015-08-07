package main

import "fmt"
import "github.com/exu/go-playground/66-godebug/dupa"

/// first run: $ go get github.com/mailgun/godebug
/// next
/// godebug run -instrument github.com/exu/go-playground/66-godebug/dupa debug.go

func main() {
	_ = "breakpoint"
	b := dupa.A()
	fmt.Println(b)
}
