package main

import "fmt"

type Cat struct {
	name  string
	skill string
}

// In Go, shallow copies of objects is something that is basically built into the language. A
// simple := assignment to a new var will create a copy as long as the target of this
// equation is not a pointer.

// However, if it is a pointer, you only need to dereference it to get at the value. Careful:
// if you do not dereference the pointer, your ‘new’ var will only be a pointer to the
// pointer and modifications done to this 'copy’ will modify the original struct.
func main() {

	cat1 := &Cat{name: "Toonces", skill: "driving"}
	cat2 := *cat1

	cat2.name = "Felix"

	fmt.Println(cat1.name, "is good at", cat1.skill)
	fmt.Println(cat2.name, "is good at", cat2.skill)

	// it'll not be copied
	cat3 := Cat{name: "Kicia", skill: "miałing"}
	cat4 := cat3

	cat2.name = "Miałcia"

	fmt.Println(cat3.name, "is good at", cat3.skill)
	fmt.Println(cat4.name, "is good at", cat4.skill)
}
