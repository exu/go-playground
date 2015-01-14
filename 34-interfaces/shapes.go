package main

import (
	"fmt"
)

type Shape interface {
	area() float64
}

type Rectangle struct {
	Width  int
	Height int
}

func (this Rectangle) area() float64 {
	return float64(this.Width) * float64(this.Height)
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	shapes := []Shape{
		Rectangle{1, 2},
		Rectangle{1, 34},
	}

	fmt.Println("Total area = ", totalArea(shapes...))
}
