package main

import (
	"fmt"
)

type Polygon interface {
	Area() int
}

type RightTriangle struct {
	c1 float64
	c2 float64
	h  float64
}

func (t *RightTriangle) Area() int {
	return int(t.c1 * t.c2 / 2)
}

func main() {
	var t Polygon
	t = &RightTriangle{3, 4, 5}
	fmt.Println(t.Area())

	var s interface{}
	s = "foo"
	x := s.(string)
	s = 5
	fmt.Println(x.(type), s.(type))
}
