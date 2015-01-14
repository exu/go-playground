package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	*Person
	SerialNumber string
	Name         string
}

func main() {
	c3po := &Android{Person: &Person{Name: "C3PO"}, SerialNumber: "hello3434"}
	droid := &Android{Person: &Person{Name: "Moto"}, SerialNumber: "hellomoto"}

	c3po.Talk()
	droid.Talk()
	c3po.Talk()
	droid.Talk()

}
