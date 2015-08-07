package main

import (
	"fmt"
	"testing"
	"testing/quick"
)

func EvenOdds(x int) string {
	if x%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func TesEvenOdds(t *testing.T) {

	f := func(x int) bool {
		y := EvenOdds(x)
		fmt.Println(x%2 == 0, y == "even", y, x)
		return x%2 == 0 && y == "even" || x%2 != 0 && y == "odd"
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
