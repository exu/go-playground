package main

import (
	"fmt"
	"github.com/speps/go-hashids"
)

func main() {
	hashData := hashids.NewData()
	hashData.Salt = "kasiainthekitchen"
	hashData.MinLength = 2
	h := hashids.NewWithData(hashData)

	e, _ := h.Encode([]int{45, 434, 1313, 99})
	fmt.Println(e)
	d := h.Decode(e)
	fmt.Println(d)

	for i := 0; i < 1000; i++ {
		e, _ = h.Encode([]int{i})
		d = h.Decode(e)
		fmt.Println(e, d)
	}

}
