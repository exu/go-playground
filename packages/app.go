package main

import "fmt"
import "./p1"
import "./p1/sub"

func main() {
	fmt.Println("Something")
	p1.Export()
	p1.Popo()
	sub.Abrakadabra()
}
