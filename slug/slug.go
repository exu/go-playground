package main

import (
	s3 "bitbucket.org/exu/kgo/text/slug"
	"fmt"
	s1 "github.com/extemporalgenome/slug"
	s2 "github.com/stvp/slug"
)

func main() {
	s := "ŁÓĄŻŚŹĘÓŁĆŃł12.łążśźęćółń"
	fmt.Println(s)
	fmt.Println(s1.SlugAscii(s))
	fmt.Println(s2.Clean(s))
	fmt.Println(s3.Slug(s))
}
